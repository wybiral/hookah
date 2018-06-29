package input

import (
	"io"
	"net/http"
	"sync"
)

type httpListenApp struct {
	server *http.Server
	// Channel of messages
	ch chan []byte
	// Lock for changing top
	mu *sync.Mutex
	// Current message being read
	top []byte
}

// HTTPListen creates an HTTP listener and return as ReadCloser
func HTTPListen(addr string) (io.ReadCloser, error) {
	app := &httpListenApp{}
	app.server = &http.Server{
		Addr:    addr,
		Handler: http.HandlerFunc(app.handle),
	}
	app.ch = make(chan []byte)
	app.mu = &sync.Mutex{}
	app.top = nil
	go app.server.ListenAndServe()
	return app, nil
}

func (app *httpListenApp) Read(b []byte) (int, error) {
	app.mu.Lock()
	defer app.mu.Unlock()
	if len(app.top) == 0 {
		app.top = <-app.ch
		if len(app.top) == 0 {
			// ch is closed
			return 0, io.EOF
		}
	}
	n := copy(b, app.top)
	app.top = app.top[n:]
	return n, nil
}

func (app *httpListenApp) Close() error {
	close(app.ch)
	return app.server.Close()
}

func (app *httpListenApp) handle(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	buffer := make([]byte, bufferSize)
	for {
		n, err := r.Body.Read(buffer)
		if err != nil {
			return
		}
		app.ch <- buffer[:n]
	}
}
