package urlshortener

import (
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
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedYaml)
	return MapHandler(pathMap, fallback), nil
}

type YAMLUrl struct {
	Path string
	Url  string
}

func parseYAML(yml []byte) ([]YAMLUrl, error) {
	var yamlUrls []YAMLUrl

	err := yaml.Unmarshal(yml, &yamlUrls)
	if err != nil {
		return nil, err
	}

	return yamlUrls, nil
}

func buildMap(yamlUrls []YAMLUrl) map[string]string {
	m := make(map[string]string)

	for _, url := range yamlUrls {
		m[url.Path] = url.Url
	}

	return m
}
