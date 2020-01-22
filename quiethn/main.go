package main

import (
	"errors"
	"io"
	"log"
	"net/http"
)

type model struct {
	url   string
	story string
}

/**
	For items: https://hacker-news.firebaseio.com/v0/item/<itemid>.json
	Forr top stories: https://hacker-news.firebaseio.com/v0/topstories.json
*/
const apiURL = "https://hacker-news.firebaseio.com/v0/"

func main() {

}

func getContent(link string) []model {
	resp, err := http.Get(link)

	if !errors.Is(err, nil) {
		log.Fatal("failed to retrieve content")
	}

	body := resp.Body
	defer body.Close()

	return hydrateContentToStruct(body)
}

func hydrateContentToStruct(reader io.ReadCloser) []model {
	// TODO: implement
}
