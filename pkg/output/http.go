package output

import (
	"io"
	"net/http"
	"net/url"
)

// HTTP creates an HTTP client and returns WriteCloser
func HTTP(addr string, opts url.Values) (io.WriteCloser, error) {
	return httprequest("http://" + addr)
}

// HTTPS creates an HTTPS client and returns WriteCloser
func HTTPS(addr string, opts url.Values) (io.WriteCloser, error) {
	return httprequest("https://" + addr)
}

func httprequest(addr string) (io.WriteCloser, error) {
	pr, pw := io.Pipe()
	req, err := http.NewRequest("PUT", addr, pr)
	if err != nil {
		return nil, err
	}
	e := make(chan error)
	go func() {
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			e <- err
			return
		}
		e <- res.Body.Close()
	}()
	cl := &httpClient{
		w: pw,
		e: e,
	}
	return cl, nil
}

type httpClient struct {
	w io.WriteCloser
	e chan error
}

func (h *httpClient) Write(b []byte) (int, error) {
	return h.w.Write(b)
}

func (h *httpClient) Close() error {
	err := h.w.Close()
	if err != nil {
		return err
	}
	return <-h.e
}
