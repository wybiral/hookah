package input

import (
	"io"
	"net/http"
	"net/url"

	"github.com/wybiral/hookah/pkg/chreader"
)

type httpListenApp struct {
	server *http.Server
	// Channel of messages
	ch chan []byte
	// ch Reader
	r io.Reader
}

// HTTPListen creates an HTTP listener and returns ReadCloser
func HTTPListen(addr string, opts url.Values) (io.ReadCloser, error) {
	app := &httpListenApp{}
	app.server = &http.Server{
		Addr:    addr,
		Handler: http.HandlerFunc(app.handle),
	}
	app.ch = make(chan []byte)
	app.r = chreader.New(app.ch)
	go app.server.ListenAndServe()
	return app, nil
}

func (app *httpListenApp) Read(b []byte) (int, error) {
	return app.r.Read(b)
}

func (app *httpListenApp) Close() error {
	// Closing ch causes r.Read to return EOF
	close(app.ch)
	return app.server.Close()
}

func (app *httpListenApp) handle(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	for {
		b := make([]byte, bufferSize)
		n, err := r.Body.Read(b)
		if err != nil && err != io.EOF {
			return
		}
		app.ch <- b[:n]
		if err == io.EOF {
			return
		}
	}
}
