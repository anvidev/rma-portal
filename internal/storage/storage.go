package storage

import (
	"context"
	"io"
)

type Storager interface {
	Put(ctx context.Context, body io.Reader, namespace, filename string) (*StorageResponse, error)
	Get()
	List(ctx context.Context) ([]StorageResponse, error)
}

type StorageResponse struct {
	Key string
	URL string
}
