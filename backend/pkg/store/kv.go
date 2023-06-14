package store

import "context"

type KV interface {
	Get(ctx context.Context, key string) (string, error)
}
