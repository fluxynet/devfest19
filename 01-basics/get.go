package main

import (
	"encoding/json"
	"net/http"
)

// getHandler handler for http get requests
func getHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

	values := r.URL.Query()

	encoder := json.NewEncoder(w)
	encoder.Encode(values)
}
