package main

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

type errorEnvelope struct {
	Error     string `json:"error"`
	RequestID string `json:"request_id"`
}

func (api *api) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	api.logger.Errorw("internal server error", "method", r.Method, "path", r.URL.Path, "request id", getRequestID(r), "error", err.Error())
	writeJSON(w, http.StatusInternalServerError, &errorEnvelope{Error: "internal server error", RequestID: getRequestID(r)})
}

func (api *api) badRequestError(w http.ResponseWriter, r *http.Request, err error) {
	api.logger.Infow("bad request error", "method", r.Method, "path", r.URL.Path, "request id", getRequestID(r), "error", err.Error())
	writeJSON(w, http.StatusBadRequest, &errorEnvelope{Error: err.Error(), RequestID: getRequestID(r)})
}

func (api *api) notFoundError(w http.ResponseWriter, r *http.Request, err error) {
	api.logger.Warnw("not found error", "method", r.Method, "path", r.URL.Path, "request id", getRequestID(r), "error", err.Error())
	writeJSON(w, http.StatusNotFound, &errorEnvelope{Error: err.Error(), RequestID: getRequestID(r)})
}

func (api *api) conflictError(w http.ResponseWriter, r *http.Request, err error) {
	api.logger.Errorw("conflict error", "method", r.Method, "path", r.URL.Path, "request id", getRequestID(r), "error", err.Error())
	writeJSON(w, http.StatusConflict, &errorEnvelope{Error: err.Error(), RequestID: getRequestID(r)})
}

func (api *api) unauthorizedError(w http.ResponseWriter, r *http.Request, err error) {
	api.logger.Warnw("unauthorized error", "method", r.Method, "path", r.URL.Path, "request id", getRequestID(r), "error", err.Error())
	writeJSON(w, http.StatusUnauthorized, &errorEnvelope{Error: err.Error(), RequestID: getRequestID(r)})
}

func (api *api) tooManyRequests(w http.ResponseWriter, r *http.Request, err error) {
	api.logger.Warnw("too many requests error", "method", r.Method, "path", r.URL.Path, "request id", getRequestID(r), "error", err.Error())
	writeJSON(w, http.StatusTooManyRequests, &errorEnvelope{Error: err.Error(), RequestID: getRequestID(r)})
}

func getRequestID(r *http.Request) string {
	rid, ok := r.Context().Value(middleware.RequestIDKey).(string)
	if !ok {
		return ""
	}
	return rid
}
