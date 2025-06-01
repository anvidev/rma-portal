package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/anvidev/rma-portal/internal/contextkeys"
	"github.com/anvidev/rma-portal/internal/store"
	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrAuthorizationMissing   = fmt.Errorf("authorization header is missing")
	ErrAuthorizationMalformed = fmt.Errorf("authorization header is malformed")
	ErrAuthorizationDenied    = fmt.Errorf("access denied")
	ErrAuthorizationExpired   = fmt.Errorf("access expired")

	ErrTooManyRequests = fmt.Errorf("too many requests")
)

func (api *api) bearerAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
			api.unauthorizedError(w, r, ErrAuthorizationMalformed)
			return
		}

		ctx := context.WithValue(r.Context(), contextkeys.UserKey, loggedInUser)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (api *api) getAuthenticatedUser(r *http.Request) (store.User, bool) {
	user, found := r.Context().Value(contextkeys.UserKey).(store.User)
	return user, found
}

func (api *api) ratelimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			api.internalServerError(w, r, err)
			return
		}

		visitor := api.baseRateLimit.Visitor(ip)

		if !visitor.Allow() {
			api.tooManyRequests(w, r, ErrTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
