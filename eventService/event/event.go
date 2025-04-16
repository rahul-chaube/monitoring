package event

import (
	"Monitoring/eventService/model"
	"Monitoring/eventService/repository"
	"fmt"
)

type EventService interface {
	AddEvent(event model.Event) (model.Event, error)
	GetEventById(id int) (model.Event, error)
	GetAllEvents() ([]model.Event, error)
	UpdateEvent(event model.Event) (model.Event, error)
	DeleteEventById(id int) error
}

type event struct {
	repository repository.EventRepository
}

func NewEventService(eventRepo repository.EventRepository) EventService {
	return event{
		repository: eventRepo,
	}
}

func (h event) AddEvent(event model.Event) (model.Event, error) {
	fmt.Println(" AddEvent 111")
	eventAdd, err := h.repository.AddEvent(event)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(eventAdd)
	return event, nil
}
func (h event) GetEventById(id int) (model.Event, error) {
	return model.Event{}, nil
}
func (h event) GetAllEvents() ([]model.Event, error) {
	events, err := h.repository.GetAllEvents()
	if err != nil {
		fmt.Println(err)
	}
	return events, nil
}
func (h event) UpdateEvent(event model.Event) (model.Event, error) {
	return event, nil
}
func (h event) DeleteEventById(id int) error {
	return nil
}
