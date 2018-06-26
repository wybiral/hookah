package output

import (
	"io"
	"net/http"
)

// Create an HTTP client and return as WriteCloser
func httpClient(addr string) (io.WriteCloser, error) {
	pr, pw := io.Pipe()
	req, err := http.NewRequest("PUT", addr, pr)
	if err != nil {
		return nil, err
	}
	go func() {
		res, _ := http.DefaultClient.Do(req)
		res.Body.Close()
		pw.Close()
	}()
	return pw, nil
}
