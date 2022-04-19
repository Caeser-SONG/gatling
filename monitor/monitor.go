package monitor

import (
	"sync/atomic"
	"time"
)

type Monitor struct {
	count
}

func InitMonitor() {

}

func (m *Monitor) Watch() {
	for range time.Tick(time.Second) {
		atomic.LoadInt32(&m.count)

		atomic.StoreInt32(&m.count, 0)
	}
}
