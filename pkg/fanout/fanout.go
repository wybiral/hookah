package fanout

import (
	"sync"
)

type Fanout struct {
	mutex *sync.RWMutex
	chans map[chan []byte]struct{}
}

func NewFanout() *Fanout {
	return &Fanout{
		mutex: &sync.RWMutex{},
		chans: make(map[chan []byte]struct{}),
	}
}

func (fan *Fanout) AddChan(ch chan []byte) {
	fan.mutex.Lock()
	defer fan.mutex.Unlock()
	fan.chans[ch] = struct{}{}
}

func (fan *Fanout) RemoveChan(ch chan []byte) {
	fan.mutex.Lock()
	defer fan.mutex.Unlock()
	delete(fan.chans, ch)
}

func (fan *Fanout) Count() int {
	fan.mutex.RLock()
	defer fan.mutex.RUnlock()
	return len(fan.chans)
}

func (fan *Fanout) Send(data []byte) {
	fan.mutex.RLock()
	defer fan.mutex.RUnlock()
	for ch, _ := range fan.chans {
		select {
		case ch <- data:
		default:
			continue
		}
	}
}
