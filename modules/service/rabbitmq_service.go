package service

import (
	"context"
	"go-web-template/modules/constant/exception"
	"go-web-template/modules/rabbitmq"
	"go-web-template/modules/util/check"
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

type IRabbitMQService interface {
	SendMessage(topic, message string) error
}

type RabbitMQService struct {
	RabbitMQManager rabbitmq.IRabbitMQManager
	Checker         check.Checker
}

func (rs RabbitMQService) SendMessage(topic string, message string) error {
	if err := rs.Checker.String(
		[]string{topic, message},
		func(str string) bool {
			return len(str) != 0
		},
	); err != nil {
		return exception.ErrInvalidData
	}

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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = channel.PublishWithContext(
		ctx,
		exchangerName,
		topic,
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		log.Printf("%s: %s", exception.ErrPublishRabbitMQMessageFailed, err)

		return exception.ErrPublishRabbitMQMessageFailed
	}

	return nil
}
