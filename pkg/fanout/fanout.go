// This package implements a fanout type that makes it easy to build
// publish-subcribe patterns using []byte channels.
package fanout

import (
	"sync"
)

// A Fanout instance manages subscribers and publishing across []byte channels.
type Fanout struct {
	// Subscribed channels
	chans map[chan []byte]struct{}
	// Mutex for chans map
	sync.RWMutex
}

// Create a new instance of Fanout.
func New() *Fanout {
	return &Fanout{
		chans: make(map[chan []byte]struct{}),
	}
}

// Add a []byte chan as a subscriber.
func (f *Fanout) Add(ch chan []byte) {
	f.Lock()
	defer f.Unlock()
	f.chans[ch] = struct{}{}
}

// Add a []byte chan from subscribers.
func (f *Fanout) Remove(ch chan []byte) {
	f.Lock()
	defer f.Unlock()
	delete(f.chans, ch)
}

// Return the current number of subscribers.
func (f *Fanout) Len() int {
	f.RLock()
	defer f.RUnlock()
	return len(f.chans)
}

// Publish a []byte message to all subscribers. If a channel can't be sent to
// it will simply be skipped.
func (f *Fanout) Send(msg []byte) {
	f.RLock()
	defer f.RUnlock()
	for ch, _ := range f.chans {
		select {
		case ch <- msg:
		default:
			continue
		}
	}
}
