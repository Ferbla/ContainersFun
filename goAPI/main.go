package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type currentTimeModel struct {
	Language string    `json:"language"`
	CT       time.Time `json:"ct"`
	Version  string    `json:"version"`
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/currentTime", getCurrentTime)

	router.Run("0.0.0.0:3000")
}

func getCurrentTime(c *gin.Context) {
	var result = currentTimeModel{
		Language: "GO",
		CT:       time.Now(),
		Version:  "0.0.2",
	}

	c.IndentedJSON(http.StatusOK, result)
}
