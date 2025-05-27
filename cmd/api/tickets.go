package main

import (
	"errors"
	"io"
	"net/http"

	"github.com/anvidev/rma-portal/internal/pdf"
	"github.com/anvidev/rma-portal/internal/queue"
	"github.com/anvidev/rma-portal/internal/store"
)

type contactPayload struct {
	Company string `json:"company" validate:"required,min=3,max=100"`
	Name    string `json:"name" validate:"required,min=3,max=100"`
	Email   string `json:"email" validate:"required,email,max=100"`
	Phone   string `json:"phone" validate:"required,max=30"`
	Street  string `json:"street" validate:"required,max=100"`
	City    string `json:"city" validate:"required,max=50"`
	Zip     string `json:"zip" validate:"required,max=20,alphanum"`
	Country string `json:"country" validate:"required,iso3166_1_alpha2"`
}

type createTicketPayload struct {
	Sender       contactPayload   `json:"sender" validate:"required"`
	Billing      contactPayload   `json:"billing" validate:"required"`
	Issue        string           `json:"issue" validate:"required,min=20,max=500"`
	Categories   []store.Category `json:"categories" validate:"required,gt=0,max=5"`
	Model        *string          `json:"model" validate:"omitempty,max=50"`
	SerialNumber *string          `json:"serial_number" validate:"omitempty,max=50"`
	Quote        string           `json:"quote" validate:"required,oneof=yes no"`
	Warranty     string           `json:"warranty" validate:"required,oneof=yes no unknown"`
}

type createTicketResponse struct {
	Ticket store.Ticket `json:"ticket"`
}

func (api *api) postCreateTicket(w http.ResponseWriter, r *http.Request) {
	var payload createTicketPayload

	if err := readJSON(w, r, &payload); err != nil {
		switch err {
		case store.ErrStatusNotImplemented:
			api.badRequestError(w, r, err)
		case store.ErrCategoryNotImplemented:
			api.badRequestError(w, r, err)
		default:
			api.internalServerError(w, r, err)
		}
		return
	}

	if err := validateJSON(&payload); err != nil {
		api.badRequestError(w, r, err)
		return
	}

	ticket := store.Ticket{
		Status:       store.CREATED,
		Categories:   payload.Categories,
		Issue:        payload.Issue,
		Model:        payload.Model,
		SerialNumber: payload.SerialNumber,
		Quote:        payload.Quote,
		Warranty:     payload.Warranty,
		Sender: store.Contact{
			Company: payload.Sender.Company,
			Name:    payload.Sender.Name,
			Email:   payload.Sender.Email,
			Phone:   payload.Sender.Phone,
			Street:  payload.Sender.Street,
			City:    payload.Sender.City,
			Zip:     payload.Sender.Zip,
			Country: payload.Sender.Country,
		},
		Billing: store.Contact{
			Company: payload.Billing.Company,
			Name:    payload.Billing.Name,
			Email:   payload.Billing.Email,
			Phone:   payload.Billing.Phone,
			Street:  payload.Billing.Street,
			City:    payload.Billing.City,
			Zip:     payload.Billing.Zip,
			Country: payload.Billing.Country,
		},
	}

	if err := api.store.Tickets.Create(r.Context(), &ticket); err != nil {
		api.internalServerError(w, r, err)
		return
	}

	api.queue.Enqueue(queue.TicketCreated, ticket)

	if err := writeJSON(w, http.StatusCreated, createTicketResponse{ticket}); err != nil {
		api.internalServerError(w, r, err)
		return
	}
}

type createLogPayload struct {
	Status          store.Status `json:"status"`
	ExternalComment string       `json:"external_comment"`
	InternalComment string       `json:"internal_comment"`
}

type createLogResponse struct {
	Log store.Log `json:"log"`
}

func (api *api) postCreateLog(w http.ResponseWriter, r *http.Request) {
	loggedInUser, _ := api.getAuthenticatedUser(r)

	ticketID := r.PathValue("id")

	var payload createLogPayload

	if err := readJSON(w, r, &payload); err != nil {
		switch err {
		case store.ErrStatusNotImplemented:
			api.badRequestError(w, r, err)
		default:
			api.internalServerError(w, r, err)
		}
		return
	}

	if err := validateJSON(&payload); err != nil {
		api.badRequestError(w, r, err)
		return
	}

	log := store.Log{
		TicketID:        ticketID,
		Status:          payload.Status,
		Initiator:       loggedInUser.Name,
		ExternalComment: payload.ExternalComment,
		InternalComment: payload.InternalComment,
	}

	if err := api.store.Tickets.CreateLog(r.Context(), &log); err != nil {
		switch err {
		case store.ErrTicketNotFound:
			api.conflictError(w, r, err)
		default:
			api.internalServerError(w, r, err)
		}
		return
	}

	if err := writeJSON(w, http.StatusCreated, createLogResponse{log}); err != nil {
		api.internalServerError(w, r, err)
		return
	}
}

type listTicketsResponse struct {
	Tickets []store.Ticket `json:"tickets"`
	Total   int            `json:"total"`
	Limit   int            `json:"limit"`
	Offset  int            `json:"offset"`
}

func (api *api) getListTickets(w http.ResponseWriter, r *http.Request) {
	var filters store.TicketFilters

	if err := filters.Parse(r); err != nil {
		api.badRequestError(w, r, err)
		return
	}

	tickets, total, err := api.store.Tickets.List(r.Context(), filters)
	if err != nil {
		api.internalServerError(w, r, err)
		return
	}

	response := listTicketsResponse{
		Tickets: tickets,
		Total:   total,
		Limit:   filters.Limit,
		Offset:  filters.Offset,
	}

	if err := writeJSON(w, http.StatusOK, response); err != nil {
		api.internalServerError(w, r, err)
		return
	}
}

type getAdminTicketResponse struct {
	Ticket store.Ticket `json:"ticket"`
}

func (api *api) getAdminTicket(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	ticket, err := api.store.Tickets.GetByID(r.Context(), id)
	if err != nil {
		switch err {
		case store.ErrTicketNotFound:
			api.notFoundError(w, r, err)
		default:
			api.internalServerError(w, r, err)
		}
		return
	}

	logs, err := api.store.Tickets.ListInternalLogs(r.Context(), id)
	if err != nil {
		api.internalServerError(w, r, err)
		return
	}

	ticket.Logs = logs

	if err := writeJSON(w, http.StatusOK, getAdminTicketResponse{*ticket}); err != nil {
		api.internalServerError(w, r, err)
		return
	}
}

func (api *api) deleteAdminTicket(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if err := api.store.Tickets.DeleteByID(r.Context(), id); err != nil {
		switch err {
		case store.ErrTicketNotFound:
			api.notFoundError(w, r, err)
		default:
			api.internalServerError(w, r, err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}

type getPublicTicketResponse struct {
	Ticket store.Ticket `json:"ticket"`
}

func (api *api) getPublicTicket(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue(("id"))

	ctx := r.Context()

	ticket, err := api.store.Tickets.GetByID(ctx, id)
	if err != nil {
		switch err {
		case store.ErrTicketNotFound:
			api.notFoundError(w, r, err)
		default:
			api.internalServerError(w, r, err)
		}
		return
	}

	logs, err := api.store.Tickets.ListExternalLogs(ctx, id)
	if err != nil {
		api.internalServerError(w, r, err)
		return
	}

	ticket.Logs = logs

	if err := writeJSON(w, http.StatusOK, getPublicTicketResponse{*ticket}); err != nil {
		api.internalServerError(w, r, err)
		return
	}
}

type getTicketStatusesResponse struct {
	Statuses []string `json:"statuses"`
}

func (api *api) getTicketStatuses(w http.ResponseWriter, r *http.Request) {
	var statuses []string

	for _, v := range store.Statuses {
		statuses = append(statuses, v)
	}

	if err := writeJSON(w, http.StatusOK, getTicketStatusesResponse{statuses}); err != nil {
		api.internalServerError(w, r, err)
		return
	}
}

type getTicketCategoriesResponse struct {
	Categories []string `json:"categories"`
}

func (api *api) getTicketCategories(w http.ResponseWriter, r *http.Request) {
	var categories []string

	for _, v := range store.Categories {
		categories = append(categories, v)
	}

	if err := writeJSON(w, http.StatusOK, getTicketCategoriesResponse{categories}); err != nil {
		api.internalServerError(w, r, err)
		return
	}
}

func (api *api) getPublicTicketLabel(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	ticket, err := api.store.Tickets.GetByID(r.Context(), id)
	if err != nil {
		switch err {
		case store.ErrTicketNotFound:
			api.notFoundError(w, r, err)
		default:
			api.internalServerError(w, r, err)
		}
		return
	}

	file, closeFile, err := pdf.TicketLabelPDF(ticket)
	if err != nil {
		api.internalServerError(w, r, err)
		return
	}
	defer closeFile()

	w.Header().Set("Content-Type", "application/pdf")
	w.WriteHeader(http.StatusOK)

	_, err = io.Copy(w, file)
	if err != nil {
		api.internalServerError(w, r, err)
		return
	}
}

type ticketFilesResponse struct {
	Files []string `json:"files"`
}

func (api *api) postTicketFiles(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if id == "" {
		api.badRequestError(w, r, errors.New("mangler et RMA sags id"))
		return
	}

	if err := r.ParseMultipartForm(16 << 20); err != nil {
		api.internalServerError(w, r, err)
		return
	}

	files, exists := r.MultipartForm.File["files"]
	if !exists || len(files) == 0 {
		api.badRequestError(w, r, errors.New("ingen filer modtaget"))
		return
	}

	var response ticketFilesResponse

	for _, fileHeader := range files {
		if fileHeader.Size > 4*1024*1024 {
			api.badRequestError(w, r, errors.New("fil er for stor: "+fileHeader.Filename))
			return
		}

		file, err := fileHeader.Open()
		if err != nil {
			api.internalServerError(w, r, err)
			return
		}
		defer file.Close()

		_, err = api.storage.Put(r.Context(), file, id, fileHeader.Filename)
		if err != nil {
			api.internalServerError(w, r, err)
			return
		}

		response.Files = append(response.Files, fileHeader.Filename)
	}

	if err := writeJSON(w, http.StatusCreated, response); err != nil {
		api.internalServerError(w, r, err)
		return
	}
}
