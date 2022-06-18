package main

import (
	"catalog/pkg/config"
	"catalog/pkg/database"
	"catalog/pkg/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func main() {

	// Logging configuration
	logger := config.NewJsonLogger()

	// Load app configuration
	appConfig, err := config.LoadAppConfig()
	if err != nil {
		logger.WithFields(logrus.Fields{
			"microservice": "Catalog",
			"author":       "Ciucur Daniel",
		}).Fatal(err)
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

	logger.WithFields(logrus.Fields{
		"microservice": "Catalog",
		"author":       "Ciucur Daniel",
	}).Info(fmt.Sprintf("Server configured starting on port %v", appConfig.ServerPort))

	logger.Fatal(server.ListenAndServe())
}
