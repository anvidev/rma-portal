package storage

import (
	"context"
	"io"
)

type Storager interface {
	Put(ctx context.Context, body io.Reader, namespace, filename string) (*StorageResponse, error)
	Get()
	PutPresign(ctx context.Context, key, contentType string, contentLength int64) (string, error)
}

type StorageResponse struct {
	Key string
	URL string
}
