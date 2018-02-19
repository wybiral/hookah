package fanout

import (
	"sync"
)

type Fanout struct {
	chans map[chan []byte]struct{}
	sync.RWMutex
}

func New() *Fanout {
	return &Fanout{
		chans: make(map[chan []byte]struct{}),
	}
}

func (f *Fanout) Add(ch chan []byte) {
	f.Lock()
	defer f.Unlock()
	f.chans[ch] = struct{}{}
}

func (f *Fanout) Remove(ch chan []byte) {
	f.Lock()
	defer f.Unlock()
	delete(f.chans, ch)
}

func (f *Fanout) Len() int {
	f.RLock()
	defer f.RUnlock()
	return len(f.chans)
}

func (f *Fanout) Send(data []byte) {
	f.RLock()
	defer f.RUnlock()
	for ch, _ := range f.chans {
		select {
		case ch <- data:
		default:
			continue
		}
	}
}
