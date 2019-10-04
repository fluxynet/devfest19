package main

import (
	"net/http"
)

const (
	name    = "https://uinames.com/api/?language=en"
	swanson = "http://ron-swanson-quotes.herokuapp.com/v2/quotes"
)

type nameResponse struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Gender  string `json:"gender"`
	Region  string `json:"region"`
}

type swansonResponse []string

func main() {

	staticHandler := http.FileServer(http.Dir("public")) // serve files from disk
	http.ListenAndServe(":8080", staticHandler)
}
