package store

import "context"

var kv KeyValue

type KeyValue interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string) error
}

func SetGlobalKeyValueStore(newKV KeyValue) {
	kv = newKV
}

func KV() KeyValue {
	return kv
}
