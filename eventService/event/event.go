package event

import (
	"Monitoring/eventService/model"
	"Monitoring/eventService/repository"
)

type EventService interface {
	AddEvent(event model.Event) (model.Event, error)
	GetEventById(id int) (model.Event, error)
	GetAllEvents() ([]model.Event, error)
	UpdateEvent(event model.Event) (model.Event, error)
	DeleteEventById(id int) error
}

type event struct {
}

func NewEventService(eventRepo repository.EventRepository) EventService {
	return event{}
}

func (h event) AddEvent(event model.Event) (model.Event, error) {

	return event, nil
}
func (h event) GetEventById(id int) (model.Event, error) {
	return model.Event{}, nil
}
func (h event) GetAllEvents() ([]model.Event, error) {
	return make([]model.Event, 0), nil
}
func (h event) UpdateEvent(event model.Event) (model.Event, error) {
	return event, nil
}
func (h event) DeleteEventById(id int) error {
	return nil
}
