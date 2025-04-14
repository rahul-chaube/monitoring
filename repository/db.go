package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func DbConnect() (mongo.Client, error) {

	return mongo.Client{}, nil
}
