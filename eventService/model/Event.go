package model

import (
	"errors"
	"time"
)

type EventType string

const (
	GenderDetection EventType = "gender_detection"
	CrowdDetection  EventType = "crowd_detection"
)

type Event struct {
	EventId   string    `json:"eventId" bson:"_id"`
	EventName string    `json:"eventName" bson:"eventName"`
	EventType EventType `json:"eventType" bson:"eventType"`
	Age       string    `json:"age" bson:"age"`
	CreatedAt time.Time `json:"createdAt"`
	Files     []string  `json:"files"`
}

func (e *Event) Validate() error {

	if e.EventId == "" {
		return errors.New("eventId is empty")
	}
	return nil
}
