package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// postHandler a demo get handler
func postHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not read body"))
		return
	}

	var msg struct {
		Name string
	}

	err = json.Unmarshal(body, &msg)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not decode body"))
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello " + msg.Name))
}
