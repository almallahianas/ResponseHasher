package hashing

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchURL(t *testing.T) {
	// prepare FetchUrl test parameters
	testServer := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path == "/success" {
					w.WriteHeader(http.StatusOK)
					w.Write([]byte("Successful Response"))
					return
				}
				// Simulate an error response
				if r.URL.Path == "/error" {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte("Internal Server Error"))
					return
				}
			},
		),
	)
	defer testServer.Close()

	testCases := []struct {
		inputURL string
		wantResp *http.Response
	}{
		// valid URL
		{inputURL: testServer.URL + "/success", wantResp: &http.Response{StatusCode: http.StatusOK}},
		// invalid URL
		{inputURL: "invalid_url", wantResp: nil},
		// valid URL with error
		{inputURL: testServer.URL + "/error", wantResp: &http.Response{StatusCode: http.StatusInternalServerError}}}

	for _, tc := range testCases {
		gotResp, _ := FetchUrl(tc.inputURL)

		if tc.wantResp == nil {
			if gotResp != nil {
				t.Errorf("FetchUrl(%s) = %v, expected nil\n", tc.inputURL, gotResp)
			}
		} else {
			if gotResp == nil {
				t.Errorf("FetchUrl(%s) = nil, expected non-nil response\n", tc.inputURL)
			} else if gotResp.StatusCode != tc.wantResp.StatusCode {
				t.Errorf("FetchUrl(%s) = %d, expected %d\n", tc.inputURL, gotResp.StatusCode, tc.wantResp.StatusCode)
			}
		}
	}
}
