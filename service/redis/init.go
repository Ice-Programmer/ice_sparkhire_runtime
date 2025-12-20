package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

const (
	Addr = "127.0.0.1:6379"
	DB   = 0
)

var (
	client *redis.Client
)

func InitRedis(ctx context.Context) error {
	client = redis.NewClient(&redis.Options{
		Addr:         "127.0.0.1:6379",
		DB:           DB,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolSize:     20,
		MinIdleConns: 5,
	})

	return client.Ping(ctx).Err()
}
