package hashing

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"net/http/httptest"
	"testing"
)

func TestMD5HashFromResponse(t *testing.T) {
	// prepare test response
	responseBody := []byte("Sample Result")
	resp := httptest.NewRecorder()
	resp.Body = bytes.NewBuffer(responseBody)

	// call MD5HashFromResponse with test parameter
	hash, err := MD5HashFromResponse(resp.Result())
	if err != nil {
		t.Fatalf("error occurred: %v", err)
	}

	// prepare expected hash and compare with the response
	expectedHash := md5.Sum(responseBody)
	expectedHashStr := hex.EncodeToString(expectedHash[:])

	if hash != expectedHashStr {
		t.Errorf("expected hash: %s, but got: %s", expectedHashStr, hash)
	}
}
