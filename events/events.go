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

func NewOrderCreatedEvent(event OrderCreateEvent) (*Event, error) {
	body, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}
	return &Event{
		ID:      uuid.NewString(),
		Topic:   topics.OrderCreated,
		Payload: body,
	}, nil
}

func NewOrderCancelledEvent(event OrderCancelledEvent) (*Event, error) {
	body, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}
	return &Event{
		ID:      uuid.NewString(),
		Topic:   topics.OrderCancelled,
		Payload: body,
	}, nil
}

type TicketCreateEvent struct {
	ID          int    `validate:"required,number"`
	Title       string `validate:"required"`
	Description string `validate:"required"`
	Price       int    `validate:"required,number"`
	UserID      string `validate:"required"`
}

func NewTicketCreateEvent(event TicketCreateEvent) (*Event, error) {
	body, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}
	return &Event{
		ID:      uuid.NewString(),
		Topic:   topics.TicketCreated,
		Payload: body,
	}, nil
}

type TicketUpdatedEvent struct {
	ID          int    `validate:"required,number"`
	Title       string `validate:"required"`
	Description string `validate:"required"`
	Price       int    `validate:"required,number"`
	UserID      string `validate:"required"`
}

func NewTicketUpdatedEvent(event TicketUpdatedEvent) (*Event, error) {
	body, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}
	return &Event{
		ID:      uuid.NewString(),
		Topic:   topics.TicketUpdated,
		Payload: body,
	}, nil
}
