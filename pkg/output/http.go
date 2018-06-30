package output

import (
	"io"
	"net/http"
)

// HTTP creates an HTTP client and returns WriteCloser
func HTTP(addr string) (io.WriteCloser, error) {
	return httprequest("http://" + addr)
}

// HTTPS creates an HTTPS client and returns WriteCloser
func HTTPS(addr string) (io.WriteCloser, error) {
	return httprequest("https://" + addr)
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
