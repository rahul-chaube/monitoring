package model

import (
	"errors"
	"time"
)

type Event struct {
	EventId   string    `json:"eventId" bson:"_id"`
	EventName string    `json:"eventName" bson:"eventName"`
	EventType string    `json:"eventType" bson:"eventType"`
	CreatedAt time.Time `json:"createdAt"`
	Files     []string  `json:"files"`
}

func (e *Event) Validate() error {

	if e.EventId == "" {
		return errors.New("eventId is empty")
	}
	return nil
}
