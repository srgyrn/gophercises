package main

import (
	"fmt"
	"net/http"

	"github.com/srgyrn/gophercises/urlshortener"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshortener.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
//	yaml := `
//- path: /urlshortener
//  url: https://github.com/gophercises/urlshortener
//- path: /urlshortener-final
//  url: https://github.com/gophercises/urlshortener/tree/solution
//`
//	yamlHandler, err := urlshortener.YAMLHandler([]byte(yaml), mapHandler)
//	if err != nil {
//		panic(err)
//	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mapHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
