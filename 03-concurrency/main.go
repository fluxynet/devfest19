package main

import (
	"encoding/json"
	"net/http"
)

const (
	nameURL    = "https://uinames.com/api/?language=en"
	swansonURL = "http://ron-swanson-quotes.herokuapp.com/v2/quotes"
)

type nameResponse struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Gender  string `json:"gender"`
	Region  string `json:"region"`
}

type swansonResponse []string

func main() {
	http.HandleFunc("/quotes", quotesHandler)
	http.HandleFunc("/stream", streamHandler)

	staticHandler := http.FileServer(http.Dir("public")) // serve files from disk
	http.Handle("/", staticHandler)

	go produce()

	http.ListenAndServe(":8080", nil)
}

// quotesHandler serves quotes
func quotesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(w)
	msg := getMessage()

	encoder.Encode(msg)
}
