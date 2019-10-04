package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// Streamer is a streaming service
type Streamer struct {
	listeners map[int]chan string
	mutex     sync.Mutex
	index     int
}

// NewStreamer is a factory function to get a new streamer service
func NewStreamer() *Streamer {
	var streamer Streamer
	streamer.listeners = make(map[int]chan string)
	return &streamer
}

// Start begins production and broadcasting
func (s *Streamer) Start() {
	for {
		msg := getMessage()
		content := fmt.Sprintf("<strong>%s </strong> %s", msg.Name, msg.Content)

		go s.Broadcast(content)

		time.Sleep(time.Second * 3)
	}
}

// Broadcast sends message to all listeners
func (s *Streamer) Broadcast(content string) {
	for _, listener := range s.listeners {
		listener <- content
	}
}

// Tunein allows someone to connect and start listening
func (s *Streamer) Tunein() (int, <-chan string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	id := s.index + 1
	s.index++

	listener := make(chan string)
	s.listeners[id] = listener
	return id, listener
}

// Tuneout allows someone to stop listening
func (s *Streamer) Tuneout(id int) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.listeners, id)
}

// streamHandler sends streaming data to the client
func streamHandler(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"status": false, "error": "streaming not supported"}`))
		return
	}

	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Connection", "keep-alive")
	flusher.Flush()

	id, msgs := stream.Tunein()
	defer stream.Tuneout(id)

	for msg := range msgs {
		log.Println(">> ", msg)
		fmt.Fprintf(w, "data: %s\n\n", msg)
		flusher.Flush()
	}
}
