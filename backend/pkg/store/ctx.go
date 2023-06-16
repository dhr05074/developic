package store

import "context"

type userCtxKey struct{}

func WithUsername(ctx context.Context, username string) context.Context {
	return context.WithValue(ctx, userCtxKey{}, username)
}

func UsernameFromContext(ctx context.Context) (string, bool) {
	username, ok := ctx.Value(userCtxKey{}).(string)
	return username, ok
}
