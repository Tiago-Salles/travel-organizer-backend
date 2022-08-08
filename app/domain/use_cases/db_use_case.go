package useCase

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func RunDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, e := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if e != nil {
		client.Disconnect(ctx)
		panic(e)
	}
	collection := client.Database("travels").Collection("travel")
	result, err := collection.InsertOne(context.Background(), GenerateTravel())
	if err != nil {
		panic(err)
	}
	dataFromDb, _e := collection.Find(ctx, bson.M{})
	if _e != nil {
		panic(e)
	}
	fmt.Print(result.InsertedID)
	fmt.Print(dataFromDb)
}
