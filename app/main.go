package main
package GenerateTravel


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



func fetchTravel(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, GenerateTravel())
}

func runApp(data any) {
	router := gin.Default()
	router.GET("/travel", fetchTravel)
	router.Run("localhost:4040")
}

func main() {
	
	runApp(dataFromDb)
}
