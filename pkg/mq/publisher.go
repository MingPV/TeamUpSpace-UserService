package mq

import (
	"encoding/json"
	"log"

	// "github.com/MingPV/EventService/internal/event/usecase"
	"github.com/streadway/amqp"
)

type MQPublisher interface {
	Publish(queue string, message any) error
}

type RabbitMQPublisher struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

type MessageEnvelope struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

func NewRabbitMQPublisher(rabbitURL string) MQPublisher {
	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		log.Fatalf("Failed to connect RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open channel: %v", err)
	}

	return &RabbitMQPublisher{
		conn: conn,
		ch:   ch,
	}
}

func (p *RabbitMQPublisher) Publish(eventType string, payload interface{}) error {

	q, err := p.ch.QueueDeclare(
		"notifications", // same as in consumer
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	envelope := MessageEnvelope{
		Type: eventType,
		Data: payload,
	}

	body, err := json.Marshal(envelope)
	if err != nil {
		return err
	}

	// body, err := json.Marshal(message)
	// if err != nil {
	// 	return err
	// }

	err = p.ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	return err
}
