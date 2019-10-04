package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var (
	listeners = make(map[int]chan string)
	m         sync.Mutex
	num       int
)

func produce() {
	for {
		msg := getMessage()

		go broadcast(msg)

		time.Sleep(time.Second * 3)
	}
}

func broadcast(msg message) {
	content := fmt.Sprintf(`<strong>%s: </strong>%s`, msg.Name, msg.Content)
	log.Println(content)
	for _, listener := range listeners {
		listener <- content
	}
}

func tunein() (int, <-chan string) {
	m.Lock()
	defer m.Unlock()
	id := num + 1
	num++

	listener := make(chan string)
	listeners[id] = listener
	return id, listener
}

func tuneout(id int) {
	m.Lock()
	defer m.Unlock()
	delete(listeners, id)
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

	id, msgs := tunein()
	defer tuneout(id)

	for msg := range msgs {
		log.Println(">> ", msg)
		fmt.Fprintf(w, "data: %s\n\n", msg)
		flusher.Flush()
	}
}
