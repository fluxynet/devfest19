package main

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

type message struct {
	Name    string
	Content string
}

func getName() (string, error) {
	client := http.Client{Timeout: time.Second * 3}
	resp, err := client.Get(nameURL)
	if err != nil {
		return "Anonymous", err
	}

	decoder := json.NewDecoder(resp.Body)
	defer resp.Body.Close()

	var name nameResponse
	err = decoder.Decode(&name)
	if err != nil {
		return "Anonymous", err
	}

	return name.Surname + ", " + name.Name, nil
}

func getQuote() (string, error) {
	client := http.Client{Timeout: time.Second * 3}
	resp, err := client.Get(swansonURL)
	if err != nil {
		return "...", err
	}

	decoder := json.NewDecoder(resp.Body)
	defer resp.Body.Close()

	var quote swansonResponse
	err = decoder.Decode(&quote)
	if err != nil {
		return "...", err
	}

	if len(quote) == 0 {
		return "...", nil
	}

	return quote[0], nil
}

func getMessage() message {
	var (
		wg  sync.WaitGroup
		msg message
	)

	wg.Add(2)

	go func() {
		defer wg.Done()
		msg.Name, _ = getName()
	}()

	go func() {
		defer wg.Done()
		msg.Content, _ = getQuote()
	}()

	wg.Wait()
	return msg
}
