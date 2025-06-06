package main

import (
	"fmt"
	"net/http"

	"github.com/anvidev/rma-portal/internal/storage"
	"github.com/anvidev/rma-portal/internal/store"
)

type postUploadRequest struct {
	Key           string `json:"key"`
	ReferenceID   string `json:"reference_id"`
	Filename      string `json:"file_name"`
	FileDomain    string `json:"file_domain"`
	ContentType   string `json:"content_type"`
	ContentLength int64  `json:"content_length"`
}

type postUploadReponse struct {
	PresignedUrl string `json:"presigned_url"`
}

func (api *api) postUpload(w http.ResponseWriter, r *http.Request) {
	var body postUploadRequest

	if err := readJSON(w, r, &body); err != nil {
		api.badRequestError(w, r, err)
		return
	}

	ctx := r.Context()

	if api.isDevelopment() {
		body.Key = fmt.Sprintf("dev/%s", body.Key)
	}

	presignedUrl, err := api.storage.PutPresign(ctx, body.Key, body.ContentType, body.ContentLength)
	if err != nil {
		api.internalServerError(w, r, err)
		return
	}

	storageUrl := fmt.Sprintf("%s/%s", storage.BucketUrl, body.Key)

	newFile := &store.File{
		FileName:    body.Filename,
		FileUrl:     storageUrl,
		FileDomain:  body.FileDomain,
		ReferenceID: body.ReferenceID,
		MimeType:    body.ContentType,
	}

	if err := api.store.Tickets.CreateFile(ctx, newFile); err != nil {
		api.internalServerError(w, r, err)
		return
	}

	if err := writeJSON(w, http.StatusOK, postUploadReponse{PresignedUrl: presignedUrl}); err != nil {
		api.internalServerError(w, r, err)
		return
	}
}

func (api *api) isDevelopment() bool {
	return api.config.server.env == "development"
}
