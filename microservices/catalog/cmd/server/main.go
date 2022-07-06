package main

import (
	"catalog/pkg/config"
	"catalog/pkg/data"
	"catalog/pkg/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func main() {

	// Logging configuration
	logger := config.NewJsonLogger()

	// Load app configuration
	appConfig, err := config.LoadAppConfig()
	if err != nil {
		logger.LogWithLevel(config.FATAL, err.Error())
	}

	// Database connect
	data.Connect(appConfig.DbConnStr)
	data.Migrate()

	controller := handlers.CatalogController{
		Db:     data.Instance,
		Logger: logger,
	}

	// Set up the router
	router := mux.NewRouter()
	router.HandleFunc("/product", controller.GetProducts).Methods(http.MethodGet)
	router.HandleFunc("/product", controller.AddProduct).Methods(http.MethodPost)
	router.HandleFunc("/product/{id}", controller.GetProductsById).Methods(http.MethodGet)

	router.HandleFunc("/seed", controller.SeedSelf).Methods(http.MethodGet)

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

	logger.LogWithLevel(config.FATAL, server.ListenAndServe().Error())
}
