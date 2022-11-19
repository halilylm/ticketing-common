package nats

import (
	"fmt"
	"github.com/halilylm/ticketing-common/events"
	"github.com/nats-io/stan.go"
)

type Publisher struct {
	conn stan.Conn
}

func NewPublisher(conn stan.Conn) *Publisher {
	return &Publisher{conn: conn}
}

// PublishMessage publish a message to a topic.
func (p *Publisher) PublishMessage(event *events.Event) error {
	// validate the topic
	if len(event.Topic) == 0 {
		return events.ErrMissingTopic
	}
	// publish the event to the topic's channel
	if _, err := p.conn.PublishAsync(event.Topic, event.Payload, nil); err != nil {
		return fmt.Errorf("error publishing message to topic: %w", err)
	}
	return nil
}
