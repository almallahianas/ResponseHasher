package hashing

import "strings"

func prepareRequestURL(inputURL string) string {
	if !strings.HasPrefix(inputURL, "http://") &&
		!strings.HasPrefix(inputURL, "https://") {
		inputURL = "http://" + inputURL
	}

	inputURL = strings.TrimSuffix(inputURL, "\n")
	inputURL = strings.TrimSuffix(inputURL, "\r")

	return inputURL
}
