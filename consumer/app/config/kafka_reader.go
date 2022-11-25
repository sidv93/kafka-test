package config

import (
	"context"
	"fmt"
	"os"

	kafka "github.com/segmentio/kafka-go"
)

var topic = os.Getenv("KAFKA_TOPIC")
var brokerAddress = os.Getenv("KAFKA_BROKER_URL")

var KafkaReader = kafka.NewReader(kafka.ReaderConfig{
	Brokers: []string{brokerAddress},
	Topic:   topic,
})

func Listen(ctx context.Context, cb func(msg string)) {
	for {
		message, err := KafkaReader.ReadMessage(ctx)
		if err != nil {
			fmt.Println("Could not read message", err)
		}

		fmt.Println("Received from Kafka: ", string(message.Value))
		cb(string(message.Value))
	}
}
