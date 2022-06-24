package main

import (
	"auth/jwt"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {

	token, err := jwt.GetJwtForUserId("1")
	if err != nil {
		fmt.Println("Failed getting JWT token")
	}
	fmt.Println(fmt.Sprintf("Got token: %v", token))
	fmt.Fprintf(w, token)
}

func checkAuthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Valid")
}

func main() {
	fmt.Printf("Auth microservice")

	router := mux.NewRouter()
	router.HandleFunc("/login", loginHandler).Methods(http.MethodGet)
	router.HandleFunc("/check", checkAuthHandler).Methods(http.MethodGet)

	server := http.Server{
		Addr:         ":8070",
		Handler:      router,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
