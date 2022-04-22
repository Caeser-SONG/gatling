package monitor

import (
	"sync/atomic"
	"testing"
)

func TestMonitor(t *testing.T) {
	m := NewMonitor()

	go m.Watch()

	for {
		atomic.AddInt32(&m.count, 1)
	}
}
