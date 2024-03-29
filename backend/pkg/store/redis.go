package store

import (
	"code-connect/schema"
	"context"
	"github.com/redis/go-redis/v9"
	"os"
	"time"
)

type Redis struct {
	redisClient *redis.Client
}

func (r *Redis) Incr(ctx context.Context, key string) (int64, error) {
	return r.redisClient.Incr(ctx, key).Result()
}

func (r *Redis) Expire(ctx context.Context, key string, expiration time.Duration) error {
	return r.redisClient.Expire(ctx, key, expiration).Err()
}

func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	return r.redisClient.Get(ctx, key).Result()
}

func (r *Redis) Set(ctx context.Context, key string, value string) error {
	return r.redisClient.Set(ctx, key, value, 0).Err()
}

func NewRedis(redisClient *redis.Client) *Redis {
	return &Redis{redisClient: redisClient}
}

func NewDefaultRedis(ctx context.Context) (*Redis, error) {
	redisAddr, ok := os.LookupEnv(schema.RedisAddrEnvKey)
	if !ok {
		return nil, schema.ErrMissingRedisAddr
	}

	redisClient := redis.NewClient(
		&redis.Options{
			Addr:     redisAddr,
			Password: "",
			DB:       0,
		},
	)

	if err := redisClient.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return NewRedis(redisClient), nil
}
