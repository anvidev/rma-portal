package main

import (
	"net/http"

	"github.com/anvidev/rma-portal/internal/store"
)

type createUserPayload struct {
	Name     string `json:"name" validate:"required,min=2,max=32"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

type createUserResponse struct {
	User store.User `json:"user"`
}

func (api *api) postCreateUser(w http.ResponseWriter, r *http.Request) {
	var payload createUserPayload

	if err := readJSON(w, r, &payload); err != nil {
		api.internalServerError(w, r, err)
		return
	}

	if err := validateJSON(&payload); err != nil {
		api.badRequestError(w, r, err)
		return
	}

	user := store.User{
		Name:  payload.Name,
		Email: payload.Email,
	}

	if err := user.Hash.Hash(payload.Password); err != nil {
		api.badRequestError(w, r, err)
		return
	}

	if err := api.store.Users.Create(r.Context(), &user); err != nil {
		switch err {
		case store.ErrUserDublicateEmail:
			api.conflictError(w, r, err)
		default:
			api.internalServerError(w, r, err)
		}
		return
	}

	if err := writeJSON(w, http.StatusCreated, createUserResponse{User: user}); err != nil {
		api.internalServerError(w, r, err)
		return
	}
}

type getListUsersResponse struct {
	Users  []store.User `json:"users"`
	Total  int          `json:"total"`
	Limit  int          `json:"limit"`
	Offset int          `json:"offset"`
}

func (api *api) getListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := api.store.Users.List(r.Context())
	if err != nil {
		api.internalServerError(w, r, err)
		return
	}

	if err := writeJSON(w, http.StatusOK, getListUsersResponse{users, 0, 0, 0}); err != nil {
		api.internalServerError(w, r, err)
		return
	}
}
