package output

import (
	"io"
	"net/http"
	"net/url"
)

// HTTP creates an HTTP client and returns WriteCloser
func HTTP(path string, args url.Values) (io.WriteCloser, error) {
	return httprequest("http://" + path)
}

// HTTPS creates an HTTPS client and returns WriteCloser
func HTTPS(path string, args url.Values) (io.WriteCloser, error) {
	return httprequest("https://" + path)
}

func httprequest(addr string) (io.WriteCloser, error) {
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
