package input

import (
	"io"
	"net/http"
	"net/url"
)

// HTTP creates a streaming HTTP client and returns ReadCloser
func HTTP(path string, args url.Values) (io.ReadCloser, error) {
	return httprequest("http://" + path)
}

// HTTPS creates a streaming HTTPS client and returns as ReadCloser
func HTTPS(path string, args url.Values) (io.ReadCloser, error) {
	return httprequest("https://" + path)
}

func httprequest(addr string) (io.ReadCloser, error) {
	req, err := http.NewRequest("GET", addr, nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}
