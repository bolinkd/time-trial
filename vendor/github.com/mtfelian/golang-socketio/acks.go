package gosocketio

import (
	"errors"
	"sync"

	"github.com/mtfelian/synced"
)

var (
	ErrorAckWaiterNotFound = errors.New("ack waiter not found")
)

// acks represents chans needed for Ack messages to work
type acks struct {
	count synced.Counter

	ackC  map[int]chan string
	ackMu sync.RWMutex
}

// nextId of ack waiter
func (a *acks) nextId() int {
	a.count.Inc()
	return a.count.Get()
}

// register new ack request waiter
func (a *acks) register(id int, ackC chan string) {
	a.ackMu.Lock()
	a.ackC[id] = ackC
	a.ackMu.Unlock()
}

// unregister a waiter by ack id that is unnecessary anymore
func (a *acks) unregister(id int) {
	a.ackMu.Lock()
	delete(a.ackC, id)
	a.ackMu.Unlock()
}

// obtain checks that waiter at given ack id exists and returns the appropriate chan
func (a *acks) obtain(id int) (chan string, error) {
	a.ackMu.RLock()
	defer a.ackMu.RUnlock()

	if ackC, ok := a.ackC[id]; ok {
		return ackC, nil
	}

	return nil, ErrorAckWaiterNotFound
}
