package config

import (
	"os"

	kafka "github.com/segmentio/kafka-go"
)

var topic = os.Getenv("KAFKA_TOPIC")
var brokerAddress = os.Getenv("KAFKA_BROKER_URL")

var KafkaWriter = kafka.NewWriter(kafka.WriterConfig{
	Brokers: []string{brokerAddress},
	Topic:   topic,
})
