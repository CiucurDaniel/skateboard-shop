package main

import (
	"encoding/json"
	"net/http"
	"shoppingcart/applogger"
	"shoppingcart/data"
	"time"
)

type ShoppingCartController struct {
	// TODO: Add datasource
	Logger *applogger.MyLogger
}

// Proposed endpoints

func (c ShoppingCartController) getCartItemsHandle(w http.ResponseWriter, r *http.Request) {
	products := data.GetItemsForUserID()
	productsJSON, _ := json.Marshal(products)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(productsJSON)
}

func (c ShoppingCartController) addItemToCartHandle(w http.ResponseWriter, r *http.Request) {
	data.AddItemToCart()

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Item added successfully."))
}

func (c ShoppingCartController) checkoutOrderForUserIdHandle(w http.ResponseWriter, r *http.Request) {

	data.CheckoutOrder("1") // TODO: In future read user id from request

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Order completed successfully."))
}

// Health endpoint

func (c ShoppingCartController) healthHandle(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Service is healthy."))
}

// Logging middleware

func (c ShoppingCartController) requestLogger(router http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		router.ServeHTTP(w, r)

		finishTime := time.Now()
		elapsedTime := finishTime.Sub(startTime)

		// call logger here
		c.Logger.LogHttpRequest(w, r, elapsedTime)
	})
}

// Auth middleware

func (c ShoppingCartController) authMiddleware(router http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// function to check headers and verify JWT token,
		// else return 404 unauthorized
		router.ServeHTTP(w, r)

	})
}
