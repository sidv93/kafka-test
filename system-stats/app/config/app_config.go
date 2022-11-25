package config

import "os"

type AppConfig struct {
	ProducerUrl string
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		ProducerUrl: os.Getenv("PRODUCER_URL"),
	}
}
