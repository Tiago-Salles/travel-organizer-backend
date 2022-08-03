package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	runApp()
}
