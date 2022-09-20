package rabbitmq

import (
	"fmt"
	"os"

	"github.com/google/wire"
)

// RabbitMQManager
var RabbitMQSet = wire.NewSet(
	wire.Bind(new(IRabbitMQManager), new(RabbitMQManager)),
	RabbitMQManagerProvider,
)

func RabbitMQManagerProvider() RabbitMQManager {
	url := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		os.Getenv("RABBIT_MQ_USERNAME"),
		os.Getenv("RABBIT_MQ_PASSWORD"),
		os.Getenv("RABBIT_MQ_HOST"),
		os.Getenv("RABBIT_MQ_PORT"),
	)

	return RabbitMQManager{
		URL: url,
	}
}
