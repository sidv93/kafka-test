package main

import (
	"fmt"

	"producer/app/handlers"

	"github.com/gin-gonic/gin"
)

const (
	serverUrl = "0.0.0.0:3000"
)

func main() {
	router := gin.New()

	router.POST("/api/producer/v1/stats", handlers.PushToKafka)

	router.Run(serverUrl)

	fmt.Println("Server running on port 3000")

}
