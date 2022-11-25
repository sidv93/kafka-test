package handlers

import (
	"encoding/json"
	"fmt"

	"producer/app/config"
	"producer/app/dto"

	"github.com/gin-gonic/gin"
	kafka "github.com/segmentio/kafka-go"
)

func PushToKafka(ctx *gin.Context) {
	// validate body
	statsDTO, err := dto.NewStatsPost(ctx)
	if err != nil {
		fmt.Println("validation error", err)
		ctx.JSON(400, gin.H{"message": "Invalid data", "status": "failure"})
		return
	}

	statsInBytes, err := json.Marshal(statsDTO)
	err = config.KafkaWriter.WriteMessages(ctx, kafka.Message{
		Key:   []byte("systemstats"),
		Value: statsInBytes,
	})
	if err != nil {
		fmt.Println("Could not write message ", err)
		ctx.JSON(500, gin.H{"message": "Could not write to kafka", "status": "failure"})
		return
	}

	fmt.Println("Message write complete")
	ctx.JSON(200, gin.H{"message": "Stats pushed to producer", "status": "success"})
}
