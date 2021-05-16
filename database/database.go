package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetClient() *mongo.Client {

	client, _ := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:farofa@mongo:27017"))

	return client

}
