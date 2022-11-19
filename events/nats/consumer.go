package nats

import (
	"encoding/json"
	"fmt"
	"github.com/halilylm/ticketing-common/events"
	"github.com/halilylm/ticketing-common/logger"
	"github.com/nats-io/stan.go"
)

// ConsumerOptions contains all the options which can be provided when reading events from a store.
type ConsumerOptions struct {
	// Limit the number of results to return
	Limit uint
	// Offset the results by this number, useful for paginated queries
	Offset uint
	// Topic name of the topic to consume
	Topic string
	// Logger is the logger for consumer
	Logging logger.Logger
	// GroupID name of the group
	Group string
	// AutoAckMode
	AutoAckMode bool
	RetryLimit  int
}

type consumer struct {
	conn stan.Conn
	opts *ConsumerOptions
}

func NewConsumer(conn stan.Conn, opts *ConsumerOptions) events.Consumer {
	return &consumer{conn: conn, opts: opts}
}

// Consume to a topic.
func (c *consumer) Consume() (<-chan *events.Event, error) {
	// validate the topic
	if len(c.opts.Topic) == 0 {
		return nil, events.ErrMissingTopic
	}

	log := c.opts.Logging

	// parse the options

	// setup the subscriber
	consumedEvents := make(chan *events.Event)
	handleMsg := func(m *stan.Msg) {
		// poison message handling
		if c.opts.RetryLimit > -1 && m.Redelivered && int(m.RedeliveryCount) > c.opts.RetryLimit {
			log.Errorf("retry limit exceed: %v", m.Sequence)
			m.Ack()
			return
		}

		// decode the message
		var evt events.Event
		if err := json.Unmarshal(m.Data, &evt); err != nil {
			log.Errorf("error decoding message: %v", err)
			// not acknowledging the message is the way to indicate an error occurred
			return
		}

		if !c.opts.AutoAckMode {
			evt.AckFunc = func() error {
				return m.Ack()
			}
			evt.NackFunc = func() error {
				return nil
			}
		}

		// push onto the channel and wait for the consumer to take the event off before we acknowledge it.
		consumedEvents <- &evt

		if !c.opts.AutoAckMode {
			return
		}
		if err := m.Ack(); err != nil {
			log.Errorf("error acknowledging message: %v", err)
		}
	}

	// set up the options
	consumerOpts := []stan.SubscriptionOption{
		stan.DurableName(c.opts.Topic),
		stan.SetManualAckMode(),
	}

	// connect the subscriber
	_, err := c.conn.QueueSubscribe(c.opts.Topic, c.opts.Group, handleMsg, consumerOpts...)
	if err != nil {
		return nil, fmt.Errorf("error subscribing to topic: %w", err)
	}

	return consumedEvents, nil
}
