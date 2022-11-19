package events

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/halilylm/ticketing-common/events/topics"
	"github.com/halilylm/ticketing-common/events/types"
)

type OrderCreateEvent struct {
	ID        int
	Status    types.OrderStatus
	UserID    string
	ExpiresAt *time.Time
	Ticket    struct {
		ID    int
		Price int
	}
}

type OrderCancelledEvent struct {
	ID     int
	Ticket struct {
		ID int
	}
}

func NewOrderCreatedEvent(event OrderCreateEvent) *Event {
	body, err := json.Marshal(event)
	if err != nil {
		return nil
	}
	return &Event{
		ID:      uuid.NewString(),
		Topic:   topics.OrderCreated,
		Payload: body,
	}
}

func NewOrderCancelledEvent(event OrderCancelledEvent) *Event {
	body, err := json.Marshal(event)
	if err != nil {
		return nil
	}
	return &Event{
		ID:      uuid.NewString(),
		Topic:   topics.OrderCancelled,
		Payload: body,
	}
}
