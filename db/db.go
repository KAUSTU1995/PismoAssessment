package db

import (
	"PismoAssessment/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"time"
)

var DB *sql.DB

func InitDB(cfg config.DatabaseConfig) {
	var err error
	dbConnectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Dbname, cfg.SSLMode,
	)

	// Retry logic for connecting to the database
	for i := 0; i < cfg.MaxRetries; i++ {
		DB, err = sql.Open("postgres", dbConnectionString)
		if err == nil {
			break
		}
		logrus.Warn("Failed to connect to the database. Retrying...")
		time.Sleep(time.Duration(cfg.RetryIntervalSeconds) * time.Second)
	}

	if err != nil {
		logrus.Fatal("Unable to connect to database:", err)
	}

	// Check the connection
	if err := DB.Ping(); err != nil {
		logrus.Fatal("Failed to ping database:", err)
	}
}
