package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/google/wire"
)

type IRedisTemplate interface {
	Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Get(key string, expiration time.Duration) *redis.StringCmd
}

type RedisTemplate struct {
	RedisManager IRedisManager
	Timeout      RedisTimeoutDuration
}

var redisTemplateWireSet = wire.NewSet(
	wire.Bind(new(IRedisTemplate), new(RedisTemplate)),
	RedisTemplateProvider,
)

type RedisTimeoutDuration time.Duration

func RedisTimeoutDurationProvider() RedisTimeoutDuration {
	return RedisTimeoutDuration(1 * time.Hour)
}

func RedisTemplateProvider(redisManager IRedisManager, timeout RedisTimeoutDuration) RedisTemplate {
	return RedisTemplate{
		RedisManager: redisManager,
		Timeout:      timeout,
	}
}

func (rt RedisTemplate) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	timeoutDuration := time.Duration(rt.Timeout)
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
	defer cancel()

	statusCmd := rt.RedisManager.Client().Set(ctx, key, value, expiration)

	return statusCmd
}

func (rt RedisTemplate) Get(key string, expiration time.Duration) *redis.StringCmd {
	timeoutDuration := time.Duration(rt.Timeout)
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
	defer cancel()

	stringCmd := rt.RedisManager.Client().Get(ctx, key)

	return stringCmd
}
