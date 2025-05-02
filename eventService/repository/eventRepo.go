package repository

import (
	"context"
	"log"
	"time"

	"github.com/rahul-chaube/monitoring/eventService/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EventRepository interface {
	AddEvent(event model.Event) (model.Event, error)
	GetEventById(id string) (model.Event, error)
	GetAllEvents() ([]model.Event, error)
	UpdateEvent(event model.Event) (model.Event, error)
	DeleteEventById(id int) error
	GetDetectionTypePercentages() ([]model.DetectionStats, error)
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
func (e *eventRepositoryImpl) GetEventById(id string) (model.Event, error) {
	ctx := context.Background()

	var event model.Event
	err := e.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&event)
	if err != nil {
		return model.Event{}, err
	}

	return event, nil
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

func (e *eventRepositoryImpl) GetDetectionTypePercentages() ([]model.DetectionStats, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	pipeline := mongo.Pipeline{
		{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$detectionType"},
			{Key: "count", Value: bson.D{{Key: "$sum", Value: 1}}},
		}}},
		{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: nil},
			{Key: "total", Value: bson.D{{Key: "$sum", Value: "$count"}}},
			{Key: "types", Value: bson.D{{Key: "$push", Value: bson.D{
				{Key: "detectionType", Value: "$_id"},
				{Key: "count", Value: "$count"},
			}}}},
		}}},
		{{Key: "$unwind", Value: "$types"}},
		{{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 0},
			{Key: "detectionType", Value: "$types.detectionType"},
			{Key: "count", Value: "$types.count"},
			{Key: "percentage", Value: bson.D{
				{Key: "$multiply", Value: bson.A{
					bson.D{{Key: "$divide", Value: bson.A{"$types.count", "$total"}}},
					100,
				}},
			}},
		}}},
	}

	cursor, err := e.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []model.DetectionStats
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}
