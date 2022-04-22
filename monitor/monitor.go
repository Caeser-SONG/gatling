package monitor

import (
	"fmt"
	"sync/atomic"
	"time"
)

// 嵌入 worker
type Monitor struct {
	count int32
	time  float32
}

func NewMonitor() *Monitor {
	return &Monitor{}
}

func (m *Monitor) Watch() {
	for range time.Tick(time.Second) {
		fmt.Printf("count =  %d \n", atomic.LoadInt32(&m.count))

		atomic.StoreInt32(&m.count, 0)
	}
}
