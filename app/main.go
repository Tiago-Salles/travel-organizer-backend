package main

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func startDataBase() {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(context, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err = client.Disconnect(context); err != nil {
		panic(err)
	}
}

type TravelModel struct {
	Country string `json:"country"`
	City    string `json:"city"`
}

func generateTravel() []TravelModel {
	var travel = []TravelModel{
		{Country: "Portugal", City: "Lisbon"},
		{Country: "Spain", City: "Madrid"},
		{Country: "Spain", City: "Barcelona"},
	}
	return travel
}

func fetchTravel(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, generateTravel())
}

func runApp() {
	router := gin.Default()
	router.GET("/travel", fetchTravel)
	router.Run("localhost:4040")
}

func main() {
	defer startDataBase()
	runApp()
}
