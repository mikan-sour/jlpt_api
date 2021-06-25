package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection
var Ctx = context.TODO()

func ConnectMongo() error {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(Ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = client.Ping(Ctx, nil)
	if err != nil {
		log.Fatal(err)
		return err
	}

	Collection = client.Database("setServiceMongo").Collection("sets")

	fmt.Println("Safely connected to MongoDB")

	return nil
}
