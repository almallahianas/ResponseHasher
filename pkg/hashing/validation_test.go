package hashing

import "testing"

func TestIsValidURL(t *testing.T) {
	// Valid URLs
	validURLs := []string{
		"https://www.example.com",
		"http://localhost:8080",
		"ftp://ftp.example.org",
	}

	// Invalid URLs
	invalidURLs := []string{
		"example.com",
		"www.example.com",
		"invalid url",
	}

	// Test valid URLs
	for _, url := range validURLs {
		if !isValidURL(url) {
			t.Errorf("expected %s to be valid, but it was considered invalid", url)
		}
	}

	// Test invalid URLs
	for _, url := range invalidURLs {
		if isValidURL(url) {
			t.Errorf("expected %s to be invalid, but it was considered valid", url)
		}
	}
}
