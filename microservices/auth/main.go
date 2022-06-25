package main

import (
	"auth/applogger"
	"auth/jwt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {

	logger := applogger.NewAppLogger()

	jwtGenerator := jwt.NewJwtGenerator([]byte("mysecretkey"))

	authController := AuthController{
		jwtGenerator: jwtGenerator,
		logger:       logger,
	}

	router := mux.NewRouter()
	router.Use(authController.requestLogger)
	router.HandleFunc("/login", authController.loginHandler)
	router.HandleFunc("/register", authController.registerHandler)

	server := http.Server{
		Addr:         ":8070",
		Handler:      router,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
