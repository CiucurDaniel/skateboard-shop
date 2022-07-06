package handlers

import (
	"catalog/pkg/config"
	"catalog/pkg/data"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

type CatalogController struct {
	Db     *gorm.DB
	Logger *config.AppLogger
}

// TODO: Implement middleware for logging

func (c CatalogController) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var products []data.Product

	result := data.Instance.Find(&products)
	if result.Error != nil {
		c.Logger.LogWithLevel(config.ERROR, fmt.Sprintf("Error occured while getting products: %v", result.Error))
		http.Error(w, errors.New("cannot retrieve products from database").Error(), http.StatusInternalServerError)
		return
	}

	productsJson, _ := json.Marshal(products)
	w.Write(productsJson)
}

func (c CatalogController) GetProductsById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product data.Product
	var id = mux.Vars(r)["id"]
	data.Instance.First(&product, "ID = ?", id)
	productJson, _ := json.Marshal(product)
	w.Write(productJson)
}

func (c CatalogController) AddProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var product data.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		c.Logger.LogWithLevel(config.ERROR, "received invalid product json")
		http.Error(w, errors.New("invalid product json").Error(), http.StatusBadRequest)
		return
	}

	result := data.Instance.Create(&product)
	if result.Error != nil {
		c.Logger.LogWithLevel(config.ERROR, fmt.Sprintf("Error occured while inserting product: %v", err))
		http.Error(w, errors.New("cannot append to database").Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Created product succesfully.")
}

// Health endpoint

func (c CatalogController) healthHandle(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Service is healthy."))
}

// TODO: DELETE LATER
func (c CatalogController) SeedSelf(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	c.Db.Create(&data.DemoProducts)
	w.Write([]byte("Inserted data please check"))
}
