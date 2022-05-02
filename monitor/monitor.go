package monitor

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 嵌入 worker
type Monitor struct {
	Count int32
	Time  int64
	Path  string
	mu    sync.Mutex
}

func NewMonitor(path string) *Monitor {
	return &Monitor{
		Path: path,
	}
}

func (m *Monitor) Watch() {
	for range time.Tick(time.Second) {
		//m.mu.Lock()
		//temp := m.Time / int64(m.Count)
		//m.mu.Unlock()

		fmt.Printf("path = %s , count =  %d , \n", m.Path, atomic.LoadInt32(&m.Count))

		atomic.StoreInt64(&m.Time, 0)
		atomic.StoreInt32(&m.Count, 0)
	}
}
