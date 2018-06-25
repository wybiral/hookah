package output

import (
	"io"
	"net/http"
	"github.com/wybiral/hookah/pkg/fanout"
)

type httpServerApp struct {
	server *http.Server
	fan *fanout.Fanout
}

// Create an HTTP server and return as WriteCloser
func httpServer(addr string) (io.WriteCloser, error) {
	app := &httpServerApp{}
	app.server = &http.Server{
		Addr: addr,
		Handler: http.HandlerFunc(app.handle),
	}
	app.fan = fanout.New()
	go app.server.ListenAndServe()
	return app, nil
}

func (app *httpServerApp) Write(b []byte) (int, error) {
	app.fan.Send(b)
	return len(b), nil
}

func (app *httpServerApp) Close() error {
	return app.server.Close()
}

func (app *httpServerApp) handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	flusher, ok := w.(http.Flusher)
	if !ok {
		return
	}
	ch := make(chan []byte, queueSize)
	app.fan.Add(ch)
	defer app.fan.Remove(ch)
	for chunk := range ch {
		_, err := w.Write(chunk)
		if err != nil {
			return
		}
		flusher.Flush()
	}
}
