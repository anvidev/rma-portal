package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/anvidev/rma-portal/internal/store"
	"github.com/golang-jwt/jwt/v5"
)

type loginRequestPayload struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type authReponse struct {
	AccessToken string     `json:"access_token"`
	User        store.User `json:"user"`
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

	// cookie := http.Cookie{
	// 	Name:     string(accessTokenKey),
	// 	Value:    token,
	// 	Expires:  time.Now().Add(time.Hour * 24),
	// 	Secure:   true,
	// 	HttpOnly: true,
	// 	SameSite: http.SameSiteDefaultMode,
	// }
	//
	// http.SetCookie(w, &cookie)

	if err := writeJSON(w, http.StatusOK, authReponse{AccessToken: token, User: user}); err != nil {
		api.internalServerError(w, r, err)
		return
	}
}

func (api *api) postValidateUser(w http.ResponseWriter, r *http.Request) {
	headertoken := r.Header.Get("Authorization")

	if headertoken == "" {
		api.unauthorizedError(w, r, ErrAuthorizationMissing)
		return
	}

	headerparts := strings.Split(headertoken, " ")

	if len(headerparts) != 2 || headerparts[0] != "Bearer" {
		api.unauthorizedError(w, r, ErrAuthorizationMalformed)
		return
	}

	token, err := api.auth.ValidateToken(headerparts[1])
	if err != nil {
		api.unauthorizedError(w, r, ErrAuthorizationDenied)
		return
	}

	exp, err := token.Claims.GetExpirationTime()
	if err != nil {
		api.unauthorizedError(w, r, ErrAuthorizationDenied)
		return
	}

	if expired := time.Now().After(exp.Time); expired {
		api.unauthorizedError(w, r, ErrAuthorizationExpired)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		api.unauthorizedError(w, r, ErrAuthorizationMalformed)
		return
	}

	sub, ok := claims["sub"]
	if !ok {
		api.unauthorizedError(w, r, ErrAuthorizationMalformed)
		return
	}

	var userID int64
	switch v := sub.(type) {
	case float64:
		userID = int64(v)
	case int:
		userID = int64(v)
	case int64:
		userID = v
	default:
		api.unauthorizedError(w, r, ErrAuthorizationMalformed)
	}

	loggedInUser, err := api.store.Users.GetByID(r.Context(), userID)
	if err != nil {
		switch err {
		case store.ErrUserNotFound:
			api.notFoundError(w, r, err)
		default:
			api.internalServerError(w, r, err)
		}
		return
	}

	newtoken, err := api.auth.GenerateToken(jwt.MapClaims{
		"iss": api.config.auth.token.host,
		"aud": api.config.auth.token.host,
		"exp": time.Now().Add(api.config.auth.token.expiry).Unix(),
		"sub": loggedInUser.ID,
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
	})
	if err != nil {
		api.internalServerError(w, r, err)
		return
	}

	if err := writeJSON(w, http.StatusOK, authReponse{AccessToken: newtoken, User: loggedInUser}); err != nil {
		api.internalServerError(w, r, err)
		return
	}
}
