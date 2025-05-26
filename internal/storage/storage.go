package storage

import (
	"context"
	"io"
)

type Storager interface {
	Put(ctx context.Context, body io.Reader, namespace, filename string) (*StorageResponse, error)
	Get()
}

type StorageResponse struct {
	Key string
	URL string
}
