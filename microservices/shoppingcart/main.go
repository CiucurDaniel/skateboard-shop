package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"shoppingcart/applogger"
	"time"
)

func main() {
	fmt.Println("Hello from shopping cart microservice")

	// logger config
	logger := applogger.NewAppLogger()
	logger.LogInfo("Logger works")

	shoppingCartController := ShoppingCartController{Logger: logger}

	router := mux.NewRouter()
	router.Use(shoppingCartController.requestLogger)
	router.HandleFunc("/health", shoppingCartController.healthHandle)
	router.HandleFunc("/checkout", shoppingCartController.checkoutOrderForUserIdHandle)

	server := http.Server{
		Addr:         ":8060",
		Handler:      router,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
