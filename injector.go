//go:build wireinject
// +build wireinject

package main

import (
	"go-web-template/modules/controller"
	"go-web-template/modules/engine"
	"go-web-template/modules/middleware"
	"go-web-template/modules/orm"
	"go-web-template/modules/rabbitmq"
	"go-web-template/modules/redis"
	"go-web-template/modules/repository"
	"go-web-template/modules/service"
	"go-web-template/modules/util"

	"github.com/google/wire"
)

func InitGinManager() *engine.GinManager {
	panic(
		wire.Build(
			engine.GinManagerModuleWireSet,
			controller.ControllerWireModuleSet,
			service.ServiceWireModuleSet,
			orm.OrmWireModuleSet,
			util.UtilWireModuleSet,
			rabbitmq.RabbitMQWireModuleSet,
			repository.RepositoryWireModuleSet,
			middleware.MiddlewareWireModuleSet,
			redis.RedisWireModuleSet,
		),
	)
}
