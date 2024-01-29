package redisctrl

import (
	"context"
	"time"

	"github.com/lwinmgmg/user-go/pkg/maths"
	"github.com/redis/go-redis/v9"
)

const (
	REDIS_DEFAULT_TIMEOUT time.Duration = time.Second
)

func getContext(defaultTimeout time.Duration, timeout ...time.Duration) (context.Context, context.CancelFunc) {
	if timeout != nil {
		defaultTimeout = maths.Sum(timeout)
	}
	return context.WithTimeout(
		context.Background(),
		defaultTimeout,
	)
}

type RedisCtrl struct {
	*redis.Client
	DefaultTimeout time.Duration
}

func (rdCtrl *RedisCtrl) GetKey(key string, timeout ...time.Duration) (string, error) {
	ctx, cancel := getContext(rdCtrl.DefaultTimeout, timeout...)
	defer cancel()
	res, err := rdCtrl.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return res, nil
}

func (rdCtrl *RedisCtrl) SetKey(key, val string, duration time.Duration, timeout ...time.Duration) error {
	ctx, cancel := getContext(rdCtrl.DefaultTimeout, timeout...)
	defer cancel()
	_, err := rdCtrl.Set(ctx, key, val, duration).Result()
	if err != nil {
		return err
	}
	return nil
}

func (rdCtrl *RedisCtrl) DelKey(keys []string, timeout ...time.Duration) error {
	ctx, cancel := getContext(rdCtrl.DefaultTimeout, timeout...)
	defer cancel()
	_, err := rdCtrl.Del(ctx, keys...).Result()
	return err
}

func NewRedisCtrl(client *redis.Client, timeout ...time.Duration) *RedisCtrl {
	defaultTimeout := REDIS_DEFAULT_TIMEOUT
	if timeout != nil {
		defaultTimeout = maths.Sum(timeout)
	}
	return &RedisCtrl{
		Client:         client,
		DefaultTimeout: defaultTimeout,
	}
}
