package main

import (
	"encoding/json"
	"errors"
	"html/template"
	"io/ioutil"
	"log"
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
	book := buildBook()

	tmpl := template.Must(template.New("index").ParseFiles([]string{"index.tmpl"}...))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		option, ok := r.URL.Query()["option"]
		data := book["intro"]

		if ok {
			data = book[option[0]]
		}

		handleError(tmpl.Execute(w, data))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func buildBook() map[string]chapter {
	file, err := ioutil.ReadFile("./adventure.json")
	handleError(err)

	var book map[string]chapter
	err = json.Unmarshal(file, &book)
	handleError(err)

	return book
}

func handleError(err error) {
	if !errors.Is(err, nil) {
		log.Fatal(err)
	}
}
