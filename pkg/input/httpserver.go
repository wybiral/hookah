package input

import (
	"io"
	"net/http"
	"sync"
)

type httpServerApp struct {
	server *http.Server
	// Channel of messages
	ch chan []byte
	// Lock for changing top
	mu *sync.Mutex
	// Current message being read
	top []byte
}

// Create an HTTP server and return as ReadCloser
func httpServer(addr string) (io.ReadCloser, error) {
	app := &httpServerApp{}
	app.server = &http.Server{
		Addr: addr,
		Handler: http.HandlerFunc(app.handle),
	}
	app.ch = make(chan []byte)
	app.mu = &sync.Mutex{}
	app.top = nil
	go app.server.ListenAndServe()
	return app, nil
}

func (app *httpServerApp) Read(b []byte) (int, error) {
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

func (app *httpServerApp) Close() error {
	close(app.ch)
	return app.server.Close()
}

func (app *httpServerApp) handle(w http.ResponseWriter, r *http.Request) {
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
