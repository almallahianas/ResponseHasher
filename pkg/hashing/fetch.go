package hashing

import (
	"fmt"
	"net/http"
)

func FetchUrl(inputURL string) (*http.Response, error) {
	requestURL := prepareRequestURL(inputURL)

	if !isValidURL(requestURL) {
		return nil, fmt.Errorf("failed to fetch URL due to invalid URL: %s", requestURL)
	}

	resp, err := http.Get(requestURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch URL due http GET err: %v ", err)
	}
	return resp, nil
}
