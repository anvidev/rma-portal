package main

import (
	"net/http"
	"time"

	"github.com/anvidev/rma-portal/internal/store"
	"github.com/golang-jwt/jwt/v5"
)

type tokenKey string

const accessTokenKey tokenKey = "access_token"

type loginRequestPayload struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type loginResponsePayload struct {
	AccessToken string `json:"access_token"`
}

func (api *api) postLoginUser(w http.ResponseWriter, r *http.Request) {
	var payload loginRequestPayload

	if err := readJSON(w, r, &payload); err != nil {
		api.internalServerError(w, r, err)
		return
	}

	if err := validateJSON(&payload); err != nil {
		api.badRequestError(w, r, err)
		return
	}

	user, err := api.store.Users.GetByEmail(r.Context(), payload.Email)
	if err != nil {
		switch err {
		case store.ErrUserNotFound:
			api.notFoundError(w, r, err)
		default:
			api.internalServerError(w, r, err)
		}
		return
	}

	if err := user.Hash.Compare(payload.Password); err != nil {
		api.unauthorizedError(w, r, err)
		return
	}

	token, err := api.auth.GenerateToken(jwt.MapClaims{
		"iss": api.config.auth.token.host,
		"aud": api.config.auth.token.host,
		"exp": time.Now().Add(api.config.auth.token.expiry).Unix(),
		"sub": user.ID,
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
	})
	if err != nil {
		api.internalServerError(w, r, err)
		return
	}

	cookie := http.Cookie{
		Name:     string(accessTokenKey),
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		Secure:   true,
		SameSite: http.SameSiteDefaultMode,
	}

	http.SetCookie(w, &cookie)

	if err := writeJSON(w, http.StatusOK, loginResponsePayload{AccessToken: token}); err != nil {
		api.internalServerError(w, r, err)
		return
	}
}

func (api *api) postLogoutUser(w http.ResponseWriter, _ *http.Request) {
	emptyCookie := http.Cookie{
		Name:   string(accessTokenKey),
		Value:  "",
		MaxAge: -1,
		Secure: true,
	}
	http.SetCookie(w, &emptyCookie)
	w.WriteHeader(http.StatusNoContent)
}
