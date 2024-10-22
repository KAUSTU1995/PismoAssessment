package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Database DatabaseConfig `json:"database"`
	Server   ServerConfig   `json:"server"`
	Logging  LoggingConfig  `json:"logging"`
}

type DatabaseConfig struct {
	Host                 string `json:"host"`
	Port                 string `json:"port"`
	User                 string `json:"user"`
	Password             string `json:"password"`
	Dbname               string `json:"dbname"`
	SSLMode              string `json:"sslmode"`
	MaxRetries           int    `json:"max_retries"`
	RetryIntervalSeconds int    `json:"retry_interval_seconds"`
}

type ServerConfig struct {
	Port string `json:"port"`
}

type LoggingConfig struct {
	Level     string `json:"level"`
	Formatter string `json:"formatter"`
}

func LoadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
