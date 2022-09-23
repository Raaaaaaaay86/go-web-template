package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/google/wire"
)

var DefaultExpiration = 12 * time.Hour

var clientInstance *redis.ClusterClient

type IRedisManager interface {
	Client() *redis.ClusterClient
}

type RedisManager struct{}

var redisManagerWireSet = wire.NewSet(
	wire.Bind(new(IRedisManager), new(RedisManager)),
	RedisManagerProvider,
)

func RedisManagerProvider() RedisManager {
	return RedisManager{}
}

func (rm RedisManager) Client() *redis.ClusterClient {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if clientInstance == nil || clientInstance.Ping(ctx).Err() != nil {
		return rm.createConnection()
	}

	return clientInstance
}

func (rm RedisManager) createConnection() *redis.ClusterClient {
	clusterClient := redis.NewClusterClient(
		&redis.ClusterOptions{
			Addrs: []string{
				"localhost:26379",
				"localhost:26380",
				"localhost:26381",
				"localhost:26382",
				"localhost:26383",
				"localhost:26384",
			},
		},
	)

	clientInstance = clusterClient

	return clientInstance
}
