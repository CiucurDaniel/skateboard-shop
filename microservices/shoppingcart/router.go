package main

import (
	"net/http"
	"shoppingcart/applogger"
	"time"
)

type ShoppingCartController struct {
	// TODO: Add datasource
	Logger *applogger.MyLogger
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
