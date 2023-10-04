package schema

import "errors"

var (
	ErrMissingRedisAddr = errors.New("missing redis addr")
)

const (
	ChatGPTAPIKeyEnvKey = "CHATGPT_API_KEY" //nolint:gosec
	RedisAddrEnvKey     = "REDIS_ADDR"
)
