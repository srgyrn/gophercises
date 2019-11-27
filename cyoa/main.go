package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"fmt"
	"net/http"
)

type options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
type chapter struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []options `json:"options"`
}

func main() {
	buildBook()

	http.ListenAndServe(":8080", func(w http.ResponseWriter, r *http.Request) http.HandlerFunc {}())
}

func buildBook() {
	file, err := ioutil.ReadFile("./adventure.json")
	handleError(err)

	bookMap := make(map[string]chapter)
	err = json.Unmarshal(file, &bookMap)
	handleError(err)

	fmt.Print(bookMap)

}

func handleError(err error) {
	if !errors.Is(err, nil) {
		log.Fatal(err)
	}
}

func handlerFunc {

}



