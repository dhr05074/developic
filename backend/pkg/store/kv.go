package store

import "context"

var kv KeyValue

type KeyValue interface {
	GetParameter(ctx context.Context, key string) (string, error)
}

func SetGlobalKeyValueStore(newKV KeyValue) {
	kv = newKV
}

func KV() KeyValue {
	return kv
}
