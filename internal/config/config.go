package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port string `envconfig:"PORT" default:"8000"`

	SqlDBHost     string `envconfig:"SQL_DB_HOST" default:"night-owl-sql-db"`
	SqlDBPort     int64  `envconfig:"SQL_DB_PORT" default:"5432"`
	SqlDBUsername string `envconfig:"SQL_DB_USERNAME" default:"root"`
	SqlDBPassword string `envconfig:"SQL_DB_PASSWORD" default:"pass"`
	SqlDBName     string `envconfig:"SQL_DB_NAME" default:"night_owl_db"`
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
