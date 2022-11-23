package events

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/halilylm/ticketing-common/events/types"
)

type OrderCreateEvent struct {
	ID        int
	Status    types.OrderStatus
	UserID    string
	ExpiresAt *time.Time
	Version   int
	Ticket    struct {
		ID    int
		Price int
	}
}

type OrderCancelledEvent struct {
	ID      int
	Version int
	Ticket  struct {
		ID int
	}
}

type TicketCreateEvent struct {
	ID          int    `validate:"required,number"`
	Title       string `validate:"required"`
	Description string `validate:"required"`
	Price       int    `validate:"required,number"`
	UserID      string `validate:"required"`
	Version     int    `validate:"required"`
}

type TicketUpdatedEvent struct {
	ID          int    `validate:"required,number"`
	Title       string `validate:"required"`
	Description string `validate:"required"`
	Price       int    `validate:"required,number"`
	UserID      string `validate:"required"`
	Version     int    `validate:"required"`
}

type OrderExpiredEvent struct {
	OrderID int
}

type PaymentCreatedEvent struct {
	ID       int    `json:"id"`
	OrderID  int    `json:"order_id"`
	StripeID string `json:"stripe_id"`
}

func NewEvent(event any, topic string) (*Event, error) {
	body, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}
	return &Event{
		ID:      uuid.NewString(),
		Topic:   topic,
		Payload: body,
	}, nil
}
