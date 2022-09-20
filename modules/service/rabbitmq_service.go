package service

import (
	"context"
	"go-web-template/modules/constant/exception"
	"go-web-template/modules/rabbitmq"
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

type IRabbitMQService interface {
	SendMessage(topic, message string) error
}

type RabbitMQService struct {
	RabbitMQManager rabbitmq.IRabbitMQManager
}

func (rs RabbitMQService) SendMessage(topic string, message string) error {
	channel := rs.RabbitMQManager.GetChannel()
	exchangerName := "random_exchanger_name"

	err := channel.ExchangeDeclare(
		exchangerName,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Printf("%s: %s", exception.ErrDeclareRabbitMQExchangerFailed, err)

		return exception.ErrDeclareRabbitMQExchangerFailed
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	err = channel.PublishWithContext(
		ctx,
		exchangerName,
		topic,
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body: []byte(message),
		},
	)
	if err != nil {
		log.Printf("%s: %s", exception.ErrPublishRabbitMQMessageFailed, err)

		return exception.ErrPublishRabbitMQMessageFailed
	}

	return nil
}
