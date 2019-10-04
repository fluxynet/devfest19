package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

// Service represents a REST service endpoint
type Service struct{}

func (s *Service) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	res := NewResponse(w)

	switch req.Method {
	case http.MethodGet:
		if req.URL.Path[len(req.URL.Path)-1] == '/' {
			s.List(res, req)
		} else {
			s.Get(res, req)
		}
	case http.MethodPost:
		s.Create(res, req)
	case http.MethodPut:
		s.Put(res, req)
	case http.MethodDelete:
		s.Delete(res, req)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	res.Send()
}

// List returns all items in the collection
func (s *Service) List(res *Response, req *http.Request) {
	res.Data = articles
}

// Get returns a single item from the collection
func (s *Service) Get(res *Response, req *http.Request) {
	id := req.URL.Path[strings.LastIndex(req.URL.Path, "/")+1 : len(req.URL.Path)]
	for i, article := range articles {
		if article.ID == id {
			res.Data = &articles[i]
			break
		}
	}

	if res.Data == nil {
		res.Status = http.StatusNotFound
		res.Data = map[string]string{"error": "item " + id + " not found"}
	}
}

// Put replaces an item in a collection
func (s *Service) Put(res *Response, req *http.Request) {
	var raw []byte

	id := req.URL.Path[strings.LastIndex(req.URL.Path, "/")+1 : len(req.URL.Path)]

	i := -1
	for j := range articles {
		if id == articles[j].ID {
			i = j
			break
		}
	}

	if i == -1 {
		res.Err = errors.New("item not found")
		res.Status = http.StatusNotFound
		return
	}

	raw, res.Err = ioutil.ReadAll(req.Body)
	if res.Err != nil {
		res.Status = http.StatusBadRequest
		return
	}

	res.Err = json.Unmarshal(raw, &articles[i])
	res.Data = articles[i]
}

// Delete removes an element from the collection
func (s *Service) Delete(res *Response, req *http.Request) {
	var p int
	id := req.URL.Path[strings.LastIndex(req.URL.Path, "/")+1 : len(req.URL.Path)]
	for i, article := range articles {
		if article.ID == id {
			res.Data = &articles[i]
			p = i
			break
		}
	}

	if res.Data == nil {
		res.Status = http.StatusNotFound
		res.Data = map[string]string{"error": "item " + id + " not found"}
		return
	}

	total := len(articles)
	articles[total-1], articles[p] = articles[p], articles[total-1] // move it to the last position
	articles = articles[:total-1]                                   // trimming off the last item is inexpensive
}

// Create adds an item to the collection
func (s *Service) Create(res *Response, req *http.Request) {
	var b Article

	decoder := json.NewDecoder(req.Body)
	res.Err = decoder.Decode(&b)
	if res.Err != nil {
		res.Status = http.StatusBadRequest
		return
	}

	articles = append(articles, b)
}
