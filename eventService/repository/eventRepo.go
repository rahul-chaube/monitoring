package repository

import "Monitoring/eventService/model"

type EventRepository interface {
	AddEvent(event model.Event) (model.Event, error)
	GetEventById(id int) (model.Event, error)
	GetAllEvents() ([]model.Event, error)
	UpdateEvent(event model.Event) (model.Event, error)
	DeleteEventById(id int) error
}

type eventRepositoryImpl struct {
}

func NewEventRepository() EventRepository {
	return &eventRepositoryImpl{}
}

func (e *eventRepositoryImpl) AddEvent(event model.Event) (model.Event, error) {
	return event, nil
}
func (e *eventRepositoryImpl) GetEventById(id int) (model.Event, error) {

	return model.Event{}, nil
}
func (e *eventRepositoryImpl) GetAllEvents() ([]model.Event, error) {

	return nil, nil
}
func (e *eventRepositoryImpl) UpdateEvent(event model.Event) (model.Event, error) {

	return event, nil
}
func (e *eventRepositoryImpl) DeleteEventById(id int) error {
	return nil
}
