// Package events is for event streaming and storage
package events

import (
	"errors"
)

var (
	// ErrMissingTopic is returned if given topic missing
	ErrMissingTopic = errors.New("missing topic")
)

// Publisher is an event publishing interface.
type Publisher interface {
	Publish(event Event) error
}

// Consumer is an event consuming interface.
type Consumer interface {
	Consume() <-chan Event
}

// Event is the object send and
// received by the broker
type Event struct {
	ID      string
	Topic   string
	Payload []byte
}
