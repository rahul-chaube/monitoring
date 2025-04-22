package repository

import (
	"context"
	"log"

	"github.com/rahul-chaube/monitoring/eventService/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EventRepository interface {
	AddEvent(event model.Event) (model.Event, error)
	GetEventById(id int) (model.Event, error)
	GetAllEvents() ([]model.Event, error)
	UpdateEvent(event model.Event) (model.Event, error)
	DeleteEventById(id int) error
}

type eventRepositoryImpl struct {
	collection *mongo.Collection
}

func NewEventRepository(db *mongo.Database) EventRepository {
	return &eventRepositoryImpl{
		collection: db.Collection("event"),
	}
}

func (e *eventRepositoryImpl) AddEvent(event model.Event) (model.Event, error) {
	event.EventId = primitive.NewObjectID().Hex()
	_, err := e.collection.InsertOne(context.Background(), event)
	if err != nil {
		log.Println(err)
		return event, err
	}
	return event, nil
}
func (e *eventRepositoryImpl) GetEventById(id int) (model.Event, error) {

	return model.Event{}, nil
}
func (e *eventRepositoryImpl) GetAllEvents() ([]model.Event, error) {
	ctx := context.Background()
	cursor, err := e.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var events []model.Event
	if err := cursor.All(ctx, &events); err != nil {
		return nil, err
	}

	return events, nil
}
func (e *eventRepositoryImpl) UpdateEvent(event model.Event) (model.Event, error) {

	return event, nil
}
func (e *eventRepositoryImpl) DeleteEventById(id int) error {
	return nil
}
