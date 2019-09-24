package gosocketio

import (
	"sync"
)

var (
	overflooded   = make(map[*Channel]struct{})
	overfloodedMu sync.Mutex
)

// CountOverfloodingChannels returns an amount of overflooding channels
func CountOverfloodingChannels() int {
	overfloodedMu.Lock()
	defer overfloodedMu.Unlock()
	return len(overflooded)
}
