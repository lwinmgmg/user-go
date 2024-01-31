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
	return rdCtrl.Get(ctx, key).Result()
}

func (rdCtrl *RedisCtrl) GetHKey(key, field string, timeout ...time.Duration) (string, error) {
	ctx, cancel := getContext(rdCtrl.DefaultTimeout, timeout...)
	defer cancel()
	return rdCtrl.HGet(ctx, key, field).Result()
}

func (rdCtrl *RedisCtrl) GetHKeyAll(key, field string, timeout ...time.Duration) (map[string]string, error) {
	ctx, cancel := getContext(rdCtrl.DefaultTimeout, timeout...)
	defer cancel()
	return rdCtrl.HGetAll(ctx, key).Result()
}

func (rdCtrl *RedisCtrl) SetKey(key, val string, duration time.Duration, timeout ...time.Duration) error {
	ctx, cancel := getContext(rdCtrl.DefaultTimeout, timeout...)
	defer cancel()
	_, err := rdCtrl.Set(ctx, key, val, duration).Result()
	return err
}

func (rdCtrl *RedisCtrl) SetHKey(key string, val any, timeout ...time.Duration) error {
	ctx, cancel := getContext(rdCtrl.DefaultTimeout, timeout...)
	defer cancel()
	_, err := rdCtrl.HSet(ctx, key, val).Result()
	return err
}

func (rdCtrl *RedisCtrl) DelKey(keys []string, timeout ...time.Duration) error {
	ctx, cancel := getContext(rdCtrl.DefaultTimeout, timeout...)
	defer cancel()
	_, err := rdCtrl.Del(ctx, keys...).Result()
	return err
}

func (rdCtrl *RedisCtrl) DelHKey(keys string, fields []string, timeout ...time.Duration) error {
	ctx, cancel := getContext(rdCtrl.DefaultTimeout, timeout...)
	defer cancel()
	_, err := rdCtrl.HDel(ctx, keys, fields...).Result()
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
