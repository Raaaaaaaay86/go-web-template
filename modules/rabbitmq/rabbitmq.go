package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"


var sharedConnection *amqp.Connection

var sharedChannel *amqp.Channel

type IRabbitMQManager interface{
	GetChannel()  *amqp.Channel
	createConnection() *amqp.Connection
	openNewChannel(conn *amqp.Connection) *amqp.Channel
}

type RabbitMQManager struct {
	URL string
}

func (rm RabbitMQManager) GetChannel() *amqp.Channel {
	if sharedChannel == nil || sharedChannel.IsClosed() {
		conn := rm.createConnection()
		channel := rm.openNewChannel(conn)
		sharedChannel = channel
	}

	return sharedChannel
}

func (rm RabbitMQManager) createConnection() *amqp.Connection {
	if sharedConnection != nil && !sharedChannel.IsClosed() {
		sharedChannel.Close()
	} 

	conn, err := amqp.Dial(rm.URL)
	if err != nil {
		panic("fail to dial RabbitMQ connection: " + err.Error())
	}

	return conn
}

func (rm RabbitMQManager) openNewChannel(conn *amqp.Connection) *amqp.Channel {
	if sharedChannel != nil && !sharedChannel.IsClosed() {
		sharedChannel.Close()
	}

	channel, err := conn.Channel()
	if err != nil {
		panic("fail to open new RabbitMQ channel: " + err.Error())
	}

	return channel
}

