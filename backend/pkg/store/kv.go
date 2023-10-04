package store

import (
	"context"
	"time"
)

var kv KeyValue

type KeyValue interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string) error
	Incr(ctx context.Context, key string) (int64, error)
	Expire(ctx context.Context, key string, expiration time.Duration) error
}

func SetGlobalKeyValueStore(newKV KeyValue) {
	kv = newKV
}

func KV() KeyValue {
	return kv
}
