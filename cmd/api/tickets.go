package main

import (
	"net/http"
	"strconv"

	"github.com/anvidev/rma-portal/internal/store"
)

type createTicketPayload struct {
	// sender info
	SenderName    string `json:"sender_name" validate:"required,min=3,max=100"`
	SenderEmail   string `json:"sender_email" validate:"required,email,max=100"`
	SenderPhone   string `json:"sender_phone" validate:"required,max=30"`
	SenderStreet  string `json:"sender_street" validate:"required,max=100"`
	SenderCity    string `json:"sender_city" validate:"required,max=50"`
	SenderZip     string `json:"sender_zip" validate:"required,max=20,alphanum"`
	SenderCountry string `json:"sender_country" validate:"required,iso3166_1_alpha2"`
	// billing info
	BillingName    string `json:"billing_name" validate:"required,min=3,max=100"`
	BillingEmail   string `json:"billing_email" validate:"required,email,max=100"`
	BillingPhone   string `json:"billing_phone" validate:"required,max=30"`
	BillingStreet  string `json:"billing_street" validate:"required,max=100"`
	BillingCity    string `json:"billing_city" validate:"required,max=50"`
	BillingZip     string `json:"billing_zip" validate:"required,max=20,alphanum"`
	BillingCountry string `json:"billing_country" validate:"required,iso3166_1_alpha2"`
	// ticket info
	Issue        string           `json:"issue" validate:"required,min=50,max=500"`
	Categories   []store.Category `json:"categories" validate:"required,gt=0,max=5"`
	Model        *string          `json:"model" validate:"omitempty,max=50"`
	SerialNumber *string          `json:"serial_number" validate:"omitempty,max=50"`
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
		Status:       store.OPEN,
		Categories:   payload.Categories,
		Issue:        payload.Issue,
		Model:        payload.Model,
		SerialNumber: payload.SerialNumber,
		Sender: store.Contact{
			Name:    payload.SenderName,
			Email:   payload.SenderEmail,
			Phone:   payload.SenderPhone,
			Street:  payload.SenderStreet,
			City:    payload.SenderCity,
			Zip:     payload.SenderZip,
			Country: payload.SenderCountry,
		},
		Billing: store.Contact{
			Name:    payload.BillingName,
			Email:   payload.BillingEmail,
			Phone:   payload.BillingPhone,
			Street:  payload.BillingStreet,
			City:    payload.BillingCity,
			Zip:     payload.BillingZip,
			Country: payload.BillingCountry,
		},
	}

	if err := api.store.Tickets.Create(r.Context(), &ticket); err != nil {
		api.internalServerError(w, r, err)
		return
	}

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

	ticketID, err := parseIntFromPath(r, "id")
	if err != nil {
		api.badRequestError(w, r, err)
	}

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
	id, err := parseIntFromPath(r, "id")
	if err != nil {
		api.badRequestError(w, r, err)
		return
	}

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
	id, err := parseIntFromPath(r, "id")
	if err != nil {
		api.badRequestError(w, r, err)
		return
	}

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
	id, err := parseIntFromPath(r, "id")
	if err != nil {
		api.badRequestError(w, r, err)
		return
	}

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

func parseIntFromPath(r *http.Request, key string) (int64, error) {
	id := r.PathValue(key)
	return strconv.ParseInt(id, 10, 64)
}
