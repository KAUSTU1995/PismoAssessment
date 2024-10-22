package main

import (
	"PismoAssessment/config"
	"PismoAssessment/controllers"
	"PismoAssessment/db"
	_ "PismoAssessment/docs"
	"PismoAssessment/middleware"
	"PismoAssessment/utils"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

// @title PismoAssessment API
// @version 1.0
// @description This is the API documentation for the Pismo Assessment project, providing endpoints for account and transaction management.
// @host localhost:8080
// @BasePath /
// @schemes http
// @contact.name Kaustubh Agarwal
// @contact.email kaustubh.agarrwal@gmail.com
func main() {

	// Load the configuration
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		logrus.Fatal("Failed to load configuration:", err)
	}

	// Set log level
	level, err := logrus.ParseLevel(cfg.Logging.Level)
	if err != nil {
		logrus.Warn("Invalid log level specified. Defaulting to info level.")
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)

	// Initialize the database
	db.InitDB(cfg.Database)

	// Initialize the validator
	utils.InitializeValidator()

	r := mux.NewRouter()

	// Use logging middleware
	r.Use(middleware.LoggingMiddleware)

	// Swagger route
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Version 1 API routes
	v1 := r.PathPrefix("/v1").Subrouter()

	// Endpoints for v1
	v1.HandleFunc("/accounts", controllers.CreateAccount).Methods("POST")
	v1.HandleFunc("/accounts/{id}", controllers.GetAccount).Methods("GET")
	v1.HandleFunc("/transactions", controllers.CreateTransaction).Methods("POST")

	// Start the server
	logrus.Infof("Starting the server on port %s", cfg.Server.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Server.Port, r))
}
