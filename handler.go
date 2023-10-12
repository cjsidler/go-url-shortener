package urlshortener

import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/yaml.v3"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		redirectUrl, found := pathsToUrls[r.URL.Path]

		if !found {
			fallback.ServeHTTP(w, r)
			return
		}

		http.Redirect(w, r, redirectUrl, http.StatusFound)
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
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.

type YAMLUrl struct {
	Path string
	Url  string
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	return func(w http.ResponseWriter, r *http.Request) {
		urls := make([]YAMLUrl, 0)

		err := yaml.Unmarshal(yml, &urls)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		fmt.Printf("--- m:\n%v\n\n", urls)

		// fmt.Println("url #1:", urls[0], urls[0].Path, urls[0].url)

		fmt.Fprintf(w, "You have reached %v", r.URL.Path)
	}, nil
}
