// Package chreader converts a []byte channel to an io.ReadCloser.
package chreader

import (
	"io"
	"sync"
)

type chReader struct {
	ch  chan []byte
	mu  sync.Mutex
	top []byte
}

// New returns an io.ReadCloser that reads from and closes ch.
func New(ch chan []byte) io.ReadCloser {
	return &chReader{ch: ch}
}

func (c *chReader) Read(b []byte) (int, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if len(c.top) == 0 {
		top := <-c.ch
		c.top = make([]byte, len(top))
		copy(c.top, top)
		if len(c.top) == 0 {
			// ch is closed
			return 0, io.EOF
		}
	}
	n := copy(b, c.top)
	c.top = c.top[n:]
	return n, nil
}

func (c *chReader) Close() error {
	close(c.ch)
	return nil
}
