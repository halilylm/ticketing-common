// Package events is for event streaming and storage
package events

import (
	"encoding/json"
	"errors"
)

var (
	// ErrMissingTopic is returned if given Topic missing
	ErrMissingTopic = errors.New("missing Topic")
	// ErrEncodingMessage is returned if there was an error encoding the message option.
	ErrEncodingMessage = errors.New("Error encoding message")
)

// Publisher is an event publishing interface.
type Publisher interface {
	PublishMessage(event *Event) error
}

// Consumer is an event consuming interface.
type Consumer interface {
	Consume() (<-chan *Event, error)
}

type AckFunc func() error

// Event is the object send and
// received by the broker
type Event struct {
	ID      string
	Topic   string
	Payload []byte
	AckFunc AckFunc
}

// Unmarshal the event into an object.
func (e *Event) Unmarshal(v any) error {
	return json.Unmarshal(e.Payload, v)
}
