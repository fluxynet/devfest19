package main

import (
	"encoding/json"
	"net/http"
)

// Response represents an http response to be sent to an http client
type Response struct {
	w      http.ResponseWriter
	Data   interface{}
	Status int
	Err    error
}

// NewResponse creates a new response object
func NewResponse(w http.ResponseWriter) *Response {
	return &Response{w: w}
}

// Send writes the response to the client
func (r *Response) Send() {
	r.w.Header().Add("Content-Type", "applicaton/json")
	encoder := json.NewEncoder(r.w)

	if r.Err == nil {
		if r.Status == 0 {
			r.Status = http.StatusOK
		}

		r.w.WriteHeader(r.Status)
		if r.Data == nil {
			r.w.Write([]byte("{}"))
		} else {
			encoder.Encode(r.Data)
		}
	} else {
		if r.Status == 0 {
			r.Status = http.StatusInternalServerError
		}
		r.w.WriteHeader(r.Status)
		encoder.Encode(map[string]string{"error": r.Err.Error()})
	}
}
