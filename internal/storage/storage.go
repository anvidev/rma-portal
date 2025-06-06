package storage

import (
	"context"
)

type Storager interface {
	PutPresign(ctx context.Context, key, contentType string, contentLength int64) (string, error)
}
