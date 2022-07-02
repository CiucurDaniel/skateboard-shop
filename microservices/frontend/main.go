package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

type frontendServer struct {
	catalogMicroservice string
	//authMicroservice    string
	//shoppingCartMicroservice string
	//orderTrackingMicroservice string
}

const port = 9000 // For the moment do not take port from env

func main() {

	fmt.Println("Welcome to Skateboard shop UI")

	svc := new(frontendServer)
	mustMapEnv(&svc.catalogMicroservice, "CATALOG_MICROSERVICE")
	//mustMapEnv(&svc.authMicroservice, "AUTH_MICROSERVICE")
	//mustMapEnv(&svc.shoppingCartMicroservice, "SHOPPING_CART_MICROSERVICE")
	//mustMapEnv(&svc.orderTrackingMicroservice, "ORDER_TRACKING_MICROSERVICE")

	r := mux.NewRouter()
	r.HandleFunc("/", svc.homeHandler).Methods(http.MethodGet, http.MethodHead)
	r.HandleFunc("/product/{id}", svc.productDetailsHandler).Methods(http.MethodGet)

	//fileServer := http.FileServer(http.Dir("./static"))
	//r.Handle("/static/", http.StripPrefix("/static/", fileServer))

	r.PathPrefix("/static/").Handler( http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	servAddr := fmt.Sprintf(":%v", port)
	server := http.Server{
		Addr:         servAddr,
		Handler:      r,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	log.Printf("starting frontend on %v", port)
	log.Fatal(server.ListenAndServe())

}

func mustMapEnv(target *string, envKey string) {
	v := os.Getenv(envKey)
	if v == "" {
		// Uncomment later, now we do not call those microservices
		//panic(fmt.Sprintf("environment variable %q not set", envKey))
	}
	*target = v
}
