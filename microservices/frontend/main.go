package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type frontendServer struct {
	// todo: here add endpoints to other microservices
	// which will be configured with env variables
}


func main(){
	svc := new(frontendServer)

	fmt.Println("Welcome to Skateboard shop UI")

	r := mux.NewRouter()
	r.HandleFunc("/", svc.homeHandler).Methods(http.MethodGet, http.MethodHead)

	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":9000", r))
}
