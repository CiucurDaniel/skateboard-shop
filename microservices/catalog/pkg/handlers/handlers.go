package handlers

import "net/http"




// Health endpoint

func healthHandle(w http.ResponseWriter, _ *http.Request){
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Service is healthy."))
}