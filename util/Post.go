package util

import (
	"io"
	"net/http"
)

func Post(url, contentType string, headers map[string]string, body io.Reader) (resp *http.Response, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	for key := range headers {
		req.Header[key] = []string{headers[key]}
	}
	return client.Do(req)
}
