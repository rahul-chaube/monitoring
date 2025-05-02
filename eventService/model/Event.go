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
	EventId       string    `json:"eventId" form:"eventId" bson:"_id"`
	Description   string    `json:"description" form:"description" bson:"description"`
	DetectionType EventType `json:"detectionType" form:"detectionType" bson:"detectionType" binding:"required,oneof=gender_detection crowd_detection"`
	Confidence    float32   `json:"confidence" form:"confidence" bson:"confidence"`
	EventTime     time.Time `json:"-" bson:"eventTime"`
	CreatedAt     time.Time `json:"-" bson:"createdAt"`
	EventTimeStr  string    `form:"eventTime" binding:"required"`
	Files         []string  `json:"files" bson:"files"`
}

func (e *Event) Validate() error {

	if e.EventId == "" {
		return errors.New("eventId is empty")
	}
	return nil
}

type DetectionStats struct {
	DetectionType string  `json:"detectionType"`
	Count         int     `json:"count"`
	Percentage    float64 `json:"percentage"`
}
