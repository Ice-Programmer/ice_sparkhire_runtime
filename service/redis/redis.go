package redis

import (
	"context"
	"errors"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
	"time"
)

func SetValue(ctx context.Context, key string, value interface{}, duration time.Duration) error {
	if v, ok := value.(string); ok {
		return client.Set(ctx, key, v, duration).Err()
	}

	bytes, err := sonic.Marshal(value)
	if err != nil {
		return err
	}
	return client.Set(ctx, key, bytes, duration).Err()
}

func GetValue(ctx context.Context, key string) (string, error) {
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func DelValue(ctx context.Context, key string) error {
	return client.Del(ctx, key).Err()
}

func HasValue(ctx context.Context, key string) bool {
	result, err := client.Exists(ctx, key).Result()
	if err != nil {
		klog.CtxErrorf(ctx, "failed to check if key exists: %s", err.Error())
		return false
	}

	klog.CtxInfof(ctx, "check if key exists: %s", result)
	return result > 0
}

func GetInterface[T any](ctx context.Context, key string) (*T, error) {
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		return nil, err
	}

	var result T
	if err := sonic.Unmarshal([]byte(val), &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func GetInterfaceList[T any](ctx context.Context, key string) ([]*T, error) {
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		return nil, err
	}

	result := make([]*T, 0)
	if err := sonic.Unmarshal([]byte(val), &result); err != nil {
		return nil, err
	}

	return result, nil
}

func GetOrSaveList[T any](ctx context.Context, key string, duration time.Duration, loader func(ctx context.Context) ([]*T, error)) ([]*T, error) {

	// 1. search from the cache
	list, err := GetInterfaceList[T](ctx, key)
	if err != nil {
		return nil, err
	}
	if len(list) > 0 {
		return list, nil
	}

	// 2. search from db
	list, err = loader(ctx)
	if err != nil {
		return nil, err
	}

	// 3. save redis value
	if len(list) > 0 {
		_ = SetValue(ctx, key, list, duration)
	}

	return list, nil
}
