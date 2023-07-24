package hashing

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
)

func MD5HashFromResponse(resp *http.Response) (string, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf(
			"failed to hash MD5 from http response, due to error reading response body with error: %v",
			err)
	}
	hash := md5.Sum(body)
	return hex.EncodeToString(hash[:]), nil
}
