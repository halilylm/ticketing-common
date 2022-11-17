package events

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
)

type KafkaPublisher struct {
	Writer  *kafka.Writer
	Topic   string
	Brokers []string
}

func NewKafkaPublisher(topic string, brokers []string) *KafkaPublisher {
	w := &kafka.Writer{
		Addr:     kafka.TCP(brokers...),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	return &KafkaPublisher{
		Writer:  w,
		Topic:   topic,
		Brokers: brokers,
	}
}

func (kp *KafkaPublisher) Publish(ctx context.Context, event Event) error {
	return kp.Writer.WriteMessages(ctx,
		kafka.Message{
			Key:   []byte(event.ID),
			Value: event.Payload,
		},
	)
}

type KafkaConsumer struct {
	Reader  *kafka.Reader
	Topic   string
	GroupID string
	Brokers []string
}

func NewKafkaConsumer(topic, groupID string, brokers []string) *KafkaConsumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	return &KafkaConsumer{
		Reader:  r,
		Topic:   topic,
		Brokers: brokers,
	}
}

func (kc *KafkaConsumer) Consume() <-chan Event {
	events := make(chan Event)
	for {
		var event Event
		m, err := kc.Reader.FetchMessage(context.Background())
		if err != nil {
			break
		}
		if err := json.Unmarshal(m.Value, &event); err != nil {
			return nil
		}
		events <- event
	}
	return events
}
