package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/srgyrn/gophercises/urlshortener"
	"github.com/srgyrn/gophercises/urlshortener/storage"
)

func main() {
	storage.InitDB()
	defer storage.Conn.CloseConnection()
	mux := defaultMux()

	// Add default data to BoltDB
	storage.Conn.AddRoute(storage.RouteData{
		Path: "/urlshort-godoc",
		Url:  "https://godoc.org/github.com/gophercises/urlshort",
	})
	storage.Conn.AddRoute(storage.RouteData{
		Path: "/yaml-godoc",
		Url:  "https://godoc.org/gopkg.in/yaml.v2",
	})

	mainHandler := urlshortener.MainHandler(mux)

	extension := flag.String("file", "json", "File extension (yaml, json)")
	flag.Parse()

	handler, err := getHandler(*extension, mainHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", handler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func getRoutesFromFile(extension string) []byte {
	filePath := "./routes/routing." + extension
	output, err := ioutil.ReadFile(filePath)

	if !errors.Is(err, nil) {
		fmt.Println("Error: " + fmt.Sprintf("failed at reading file at %s", filePath))
		os.Exit(1)
	}

	return output
}

func getHandler(fileType string, fallback http.HandlerFunc) (http.HandlerFunc, error) {
	routes := getRoutesFromFile(fileType)
	switch fileType {
	case "yaml":
		// Build the YAMLHandler using the mapHandler as the
		// fallback
		return urlshortener.YAMLHandler(routes, fallback)
	case "json":
		return urlshortener.JSONHandler(routes, fallback)
	default:
		return nil, errors.New("handler type not found")
	}
}
