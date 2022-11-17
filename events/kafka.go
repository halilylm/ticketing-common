package events

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
)

type kafkaPublisher struct {
	writer  *kafka.Writer
	topic   string
	brokers []string
}

func NewKafkaPublisher(topic string, brokers []string) *kafkaPublisher {
	w := &kafka.Writer{
		Addr:     kafka.TCP(brokers...),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	return &kafkaPublisher{
		writer:  w,
		topic:   topic,
		brokers: brokers,
	}
}

func (kp *kafkaPublisher) Publish(ctx context.Context, event Event) error {
	return kp.writer.WriteMessages(ctx,
		kafka.Message{
			Key:   []byte(event.ID),
			Value: event.Payload,
		},
	)
}

type kafkaConsumer struct {
	reader  *kafka.Reader
	topic   string
	groupID string
	brokers []string
}

func NewKafkaConsumer(topic, groupID string, brokers []string) *kafkaConsumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	return &kafkaConsumer{
		reader:  r,
		topic:   topic,
		brokers: brokers,
	}
}

func (kc *kafkaConsumer) Consume() <-chan Event {
	events := make(chan Event)
	for {
		var event Event
		m, err := kc.reader.FetchMessage(context.Background())
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
