package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection
var Ctx context.Context

func ConnectMongo() error {

	Ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

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
