package redis

import "github.com/google/wire"

var RedisWireModuleSet = wire.NewSet(
	redisTemplateWireSet,
	redisManagerWireSet,
	RedisTimeoutDurationProvider,
)
