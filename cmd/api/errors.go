package main

import "net/http"

type errorEnvelope struct {
	Error string `json:"error"`
}

func (api *api) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	api.logger.Errorw("internal server error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJSON(w, http.StatusInternalServerError, &errorEnvelope{Error: "internal server error"})
}

func (api *api) badRequestError(w http.ResponseWriter, r *http.Request, err error) {
	api.logger.Infow("bad request error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJSON(w, http.StatusBadRequest, &errorEnvelope{Error: err.Error()})
}

func (api *api) notFoundError(w http.ResponseWriter, r *http.Request, err error) {
	api.logger.Warnw("not found error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJSON(w, http.StatusNotFound, &errorEnvelope{Error: err.Error()})
}

func (api *api) conflictError(w http.ResponseWriter, r *http.Request, err error) {
	api.logger.Errorw("conflict error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJSON(w, http.StatusConflict, &errorEnvelope{Error: err.Error()})
}

func (api *api) unauthorizedError(w http.ResponseWriter, r *http.Request, err error) {
	api.logger.Warnw("unauthorized error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJSON(w, http.StatusUnauthorized, &errorEnvelope{Error: err.Error()})
}
