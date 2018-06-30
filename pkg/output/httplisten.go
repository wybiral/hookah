package output

import (
	"io"
	"net/http"
	"net/url"

	"github.com/wybiral/hookah/pkg/fanout"
)

type httpListenApp struct {
	server *http.Server
	fan    *fanout.Fanout
}

// HTTPListen creates an HTTP listener and returns WriteCloser
func HTTPListen(path string, args url.Values) (io.WriteCloser, error) {
	app := &httpListenApp{}
	app.server = &http.Server{
		Addr:    path,
		Handler: http.HandlerFunc(app.handle),
	}
	app.fan = fanout.New()
	go app.server.ListenAndServe()
	return app, nil
}

func (app *httpListenApp) Write(b []byte) (int, error) {
	app.fan.Send(b)
	return len(b), nil
}

func (app *httpListenApp) Close() error {
	return app.server.Close()
}

func (app *httpListenApp) handle(w http.ResponseWriter, r *http.Request) {
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
