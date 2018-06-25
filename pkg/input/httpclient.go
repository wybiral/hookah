package input

import (
	"io"
	"net/http"
)

// Create an HTTP client and return as ReadCloser
func httpClient(addr string) (io.ReadCloser, error) {
	req, err := http.NewRequest("GET", "http://" + addr, nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}
