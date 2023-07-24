package hashing

import (
	"net/url"
)

func isValidURL(inputURL string) bool {
	u, err := url.Parse(inputURL)
	if err != nil {
		return false
	}
	return u.Scheme != "" && u.Host != ""
}
