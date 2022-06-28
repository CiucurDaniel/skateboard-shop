package handlers

import (
	"catalog/pkg/config"
	"catalog/pkg/database"
	"catalog/pkg/model"
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
)

type CatalogController struct {
	db     *gorm.DB
	logger *config.AppLogger
}

// TODO: Implement middleware for logging

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var products []model.Product
	database.Instance.Find(&products)
	productsJson, _ := json.Marshal(products)
	w.Write(productsJson)
}

// Health endpoint

func healthHandle(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Service is healthy."))
}
