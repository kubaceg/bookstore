package eventbus

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

type RabbitMQPublisher struct {
	channel  *amqp.Channel
	exchange string
}

func NewRabbitMqPublisher(channel *amqp.Channel, exchange string) Publisher {
	return &RabbitMQPublisher{channel: channel, exchange: exchange}
}

func (r *RabbitMQPublisher) Publish(msg Message) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	return r.channel.Publish(r.exchange, msg.GetTopic(), false, false, amqp.Publishing{Body: body})
}
