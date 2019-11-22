package urlshortener

import (
	"encoding/json"
	"errors"
	"github.com/go-yaml/yaml"
	"net/http"
)

type routeData struct {
	Path string `yaml:"path" json:"path"`
	Url string `yaml:"url" json:"url"`
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if url, exists := pathsToUrls[path]; exists {
			http.Redirect(w, r, url, http.StatusFound)
		}

		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var routes []routeData
	err := yaml.Unmarshal(yml, &routes)

	if !errors.Is(err, nil) {
		return nil, err
	}

	return MapHandler(buildMap(routes), fallback), nil
}

// JSONHandler will parse the provided JSON and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// JSON is expected to be in the format:
//
//     "path": "/some-path"
//      "url": "https://www.some-url.com/demo"
//
// The only errors that can be returned all related to having
// invalid JSON data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func JSONHandler(jsn []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var routes []routeData
	err := json.Unmarshal(jsn, &routes)

	if !errors.Is(err, nil) {
		return nil, err
	}

	return MapHandler(buildMap(routes), fallback), nil
}


// buildMap takes a slice of routeData and converts it to a map
func buildMap(routes []routeData) map[string]string {
	routeMap := make(map[string]string)

	for _, line := range routes {
		routeMap[line.Path] = line.Url
	}

	return routeMap
}
