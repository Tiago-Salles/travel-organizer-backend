package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TravelModel struct {
	Country string `json:"country"`
	City    string `json:"city"`
}

func generateTravel() TravelModel {
	var travel = TravelModel{
		Country: "Portugal",
		City:    "Lisbon",
	}
	return travel
}

func fetchTravel(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, generateTravel())
}

func runApp(data any) {
	router := gin.Default()
	router.GET("/travel", fetchTravel)
	router.Run("localhost:4040")
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, e := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if e != nil {
		client.Disconnect(ctx)
		panic(e)
	}
	collection := client.Database("travels").Collection("travel")
	result, err := collection.InsertOne(context.Background(), generateTravel())
	if err != nil {
		panic(err)
	}
	dataFromDb, _e := collection.Find(ctx, bson.M{})
	if _e != nil {
		panic(e)
	}
	fmt.Print(result.InsertedID)
	runApp(dataFromDb)
}
