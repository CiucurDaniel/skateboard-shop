package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func InitializeRouter(router *mux.Router) {
	router.HandleFunc("/health", healthHandle).Methods(http.MethodGet)
}
