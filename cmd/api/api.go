package main

import (
	"net/http"
	"time"

	"github.com/anvidev/rma-portal/internal/auth"
	"github.com/anvidev/rma-portal/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

type api struct {
	logger *zap.SugaredLogger
	config config
	store  store.Store
	auth   auth.Authenticator
}

type config struct {
	server   serverConfig
	database databaseConfig
	auth     authConfig
}

type serverConfig struct {
	env          string
	addr         string
	readTimeout  time.Duration
	writeTimeout time.Duration
	idleTimeout  time.Duration
}

type databaseConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  time.Duration
}

type authConfig struct {
	token tokenConfig
}

type tokenConfig struct {
	host   string
	secret string
	expiry time.Duration
}

func (api *api) mount() http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.StripSlashes)

	r.Route("/v1", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/login", api.postLoginUser)
			r.Get("/validate", api.postValidateUser)
		})

		r.Route("/admin", func(r chi.Router) {
			r.Use(api.bearerAuthorization)
			r.Route("/users", func(r chi.Router) {
				r.Get("/", api.getListUsers)
				r.Post("/", api.postCreateUser)
				r.Route("/{userID}", func(r chi.Router) {
					r.Get("/", nil)
					r.Put("/", nil)
					r.Delete("/", nil)
				})
			})
			r.Route("/tickets", func(r chi.Router) {
				r.Get("/", api.getListTickets)
				r.Post("/", api.postCreateTicket)
				r.Route("/{id}", func(r chi.Router) {
					r.Get("/", api.getAdminTicket)
					r.Put("/", nil)
					r.Delete("/", api.deleteAdminTicket)
					r.Post("/log", api.postCreateLog)
				})
			})
		})

		r.Route("/tickets", func(r chi.Router) {
			r.Get("/statuses", api.getTicketStatuses)
			r.Get("/categories", api.getTicketCategories)
			r.Get("/{id}", api.getPublicTicket)
		})
	})

	return r
}

func (api *api) run(h http.Handler) error {
	s := &http.Server{
		Addr:         api.config.server.addr,
		ReadTimeout:  api.config.server.readTimeout,
		WriteTimeout: api.config.server.writeTimeout,
		IdleTimeout:  api.config.server.idleTimeout,
		Handler:      h,
	}

	api.logger.Infow("Server started successfully", "port", api.config.server.addr, "env", api.config.server.env)

	return s.ListenAndServe()
}
