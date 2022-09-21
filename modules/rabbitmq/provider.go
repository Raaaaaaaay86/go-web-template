package rabbitmq

import "github.com/google/wire"

var RabbitMQWireModuleSet = wire.NewSet(
	RabbitMQSet,
)
