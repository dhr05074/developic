package store

import "context"

type (
	userCtxKey struct{}
	ipCtxKey   struct{}
)

func CtxWithUsername(ctx context.Context, username string) context.Context {
	return context.WithValue(ctx, userCtxKey{}, username)
}

func UsernameFromContext(ctx context.Context) (string, bool) {
	username, ok := ctx.Value(userCtxKey{}).(string)
	return username, ok
}

func CtxWithIP(ctx context.Context, ip string) context.Context {
	return context.WithValue(ctx, ipCtxKey{}, ip)
}

func IPFromContext(ctx context.Context) (string, bool) {
	ip, ok := ctx.Value(ipCtxKey{}).(string)
	return ip, ok
}
