package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port string `envconfig:"PORT" default:"8000"`
}

func NewConfig() (*Config, error) {
	var c Config

	_ = godotenv.Load(".env")

	err := envconfig.Process("APP", &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
