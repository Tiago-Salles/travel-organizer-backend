package useCase

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func fetchTravel(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, GenerateTravel())
}

func RunApp() {
	router := gin.Default()
	router.GET("/travel", fetchTravel)
	router.Run("localhost:4040")
}
