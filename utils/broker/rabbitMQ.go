package broker

import (
	"OrderDelayServing/internal/config"
	"fmt"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	Queue   amqp.Queue
}

func InitRabbitMQ(conf *config.Config) (*RabbitMQ, error) {
	url := fmt.Sprintf("amqp://%s:%s@%s:%d/", conf.RabbitMQ.User, conf.RabbitMQ.Password,
		conf.RabbitMQ.Host, conf.RabbitMQ.Port)

	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	//defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	//defer ch.Close()

	queue, err := ch.QueueDeclare(
		conf.RabbitMQ.Queue,
		false,
		false,
		false,
		false,
		nil,
	)

	return &RabbitMQ{
		Conn:    conn,
		Channel: ch,
		Queue:   queue,
	}, nil
}
