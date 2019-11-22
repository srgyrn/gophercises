package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

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

	extension := flag.String("file", "json", "File extension (yaml, json)")
	flag.Parse()


	handler, err := getHandler(*extension, mapHandler)
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
	filePath := "./routes/routing."+extension
	file, err := os.Open(filePath)
	defer file.Close()
	checkErrorAndExit(err, "file not found | " + filePath)

	output, err := ioutil.ReadAll(bufio.NewReader(file))
	checkErrorAndExit(err, "failed at reading file")

	return output
}

func checkErrorAndExit(err error, message string) {
	if !errors.Is(err, nil) {
		fmt.Println("Error: "+message)
		os.Exit(1)
	}
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