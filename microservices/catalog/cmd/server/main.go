package main

import (
	"catalog/pkg/config"
	"catalog/pkg/database"
	"catalog/pkg/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {

	// Logging configuration
	logger := config.NewJsonLogger()

	// Load app configuration
	appConfig, err := config.LoadAppConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Database connect
	database.Connect(appConfig.DbConnStr)
	database.Migrate()

	// Set up the router
	router := mux.NewRouter()
	handlers.InitializeRouter(router)

	servAddr := fmt.Sprintf(":%v", appConfig.ServerPort)
	server := http.Server{
		Addr:         servAddr,
		Handler:      router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	logger.LogWithLevel(config.INFO, fmt.Sprintf("Server configured on port %v", appConfig.ServerPort))
	logger.LogWithLevel(config.INFO, fmt.Sprintf("Database configured on conn string %v", appConfig.DbConnStr))

	log.Fatal(server.ListenAndServe())
}
