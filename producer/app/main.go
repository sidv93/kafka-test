package main

import (
	"encoding/json"
	"fmt"

	"producer/app/dto"

	"github.com/gin-gonic/gin"
	kafka "github.com/segmentio/kafka-go"
)

const (
	serverPort    = "0.0.0.0:5000"
	topic         = "systemstats"
	brokerAddress = "broker:29092"
)

var kafkaWriter = kafka.NewWriter(kafka.WriterConfig{
	Brokers: []string{brokerAddress},
	Topic:   topic,
})

func check(c *gin.Context) {
	c.JSON(200, gin.H{"status": "healthy"})
}

func pushToKafka(ctx *gin.Context) {
	statsDTO, err := dto.NewStatsPost(ctx)
	fmt.Println("stats", statsDTO)
	if err != nil {
		fmt.Println("validation error", err)
		ctx.JSON(400, gin.H{"message": "Invalid data", "status": "failure"})
		return
	}
	statsInBytes, err := json.Marshal(statsDTO)
	err = kafkaWriter.WriteMessages(ctx, kafka.Message{
		Key:   []byte("systemStats"),
		Value: statsInBytes,
	})
	if err != nil {
		fmt.Println("Could not write message ", err)
		ctx.JSON(500, gin.H{"message": "Could not write to kafka", "status": "failure"})
		return
	}

	fmt.Println("Messafe write complete")
	ctx.JSON(200, gin.H{"message": "Stats pushed to producer", "status": "success"})
}

func main() {
	// createKafkaWriter()
	fmt.Println("Created Kafka writer")
	router := gin.New()

	router.GET("/", check)
	router.POST("/api/producer/v1/stats", pushToKafka)

	router.Run(serverPort)

	fmt.Println("Server running on port 5000")

}
