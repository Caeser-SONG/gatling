package monitor

import (
	"fmt"
	"sync/atomic"
	"time"
)

// 嵌入 worker
type Monitor struct {
	Count int32
	Time  float32
}

func NewMonitor() *Monitor {
	return &Monitor{}
}

func (m *Monitor) Watch() {
	for range time.Tick(time.Second) {
		fmt.Printf("count =  %d \n", atomic.LoadInt32(&m.Count))

		atomic.StoreInt32(&m.Count, 0)
	}
}
