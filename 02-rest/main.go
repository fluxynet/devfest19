package main

import (
	"fmt"
	"net/http"

	"github.com/bxcodec/faker"
)

// Article represents a news bulletin
type Article struct {
	ID      string `json:"id" faker:"uuid_hyphenated"`
	Title   string `json:"title" faker:"sentence"`
	Content string `json:"content" faker:"paragraph"`
	Author  Author `json:"author"`
}

// Author represents an author
type Author struct {
	ID    string `json:"id" faker:"uuid_hyphenated"`
	Name  string `json:"name" faker:"name"`
	Email string `json:"email" faker:"email"`
}

var (
	articles = []Article{} // hold values that will be served
)

func main() {
	svc := &Service{}
	http.ListenAndServe(":8080", svc)
}

func init() {
	for i := 0; i < 10; i++ {
		article := Article{}
		err := faker.FakeData(&article)

		if err == nil {
			articles = append(articles, article)
		} else {
			fmt.Println(err)
		}
	}
}
