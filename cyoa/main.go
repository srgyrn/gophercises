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

type storyData struct {
	book map[string]chapter
	tpl *template.Template
}

func main() {
	tmpl := template.Must(template.ParseFiles([]string{"index.tmpl"}...))
	defaultHandler := defaultMux(storyData{buildBook(), tmpl})

	log.Fatal(http.ListenAndServe(":8080", defaultHandler))
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

func defaultMux(storyData storyData) http.Handler {
	mux := http.NewServeMux()
	handler := 	func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			option, ok := r.URL.Query()["option"]
			data := storyData.book["intro"]

			if ok {
				data = storyData.book[option[0]]
			}

			handleError(storyData.tpl.Execute(w, data))
		}
	}()
	mux.Handle("/", handler)

	return mux
}