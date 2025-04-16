package model

import "time"

type Event struct {
	EventId   string    `json:"eventId"`
	EventName string    `json:"eventName"`
	EventType string    `json:"eventType"`
	CreatedAt time.Time `json:"createdAt"`
	Files     []string  `json:"files"`
}
