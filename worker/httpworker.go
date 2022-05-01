package worker

import (
	"gatling/client"
	"gatling/monitor"
	"sync/atomic"
)

type HttpWorker struct {
	input   HttpInput
	cancel  chan struct{}
	client  *client.HttpClient
	monitor *monitor.Monitor
}

type HttpInput struct {
	Method string
	Url    string
	Header map[string]string
	Data   map[string]interface{}
}

func NewHttpWorker(input HttpInput, monitor *monitor.Monitor) *HttpWorker {
	return &HttpWorker{
		input:   input,
		monitor: monitor,
		cancel:  make(chan struct{}),
		client:  client.NewHttpClient(input.Method, input.Url, input.Data, input.Header),
	}
}

// 停止一个并发
func (w *HttpWorker) Cancel() {
	w.cancel <- struct{}{}
}

// go  run this function
func (w *HttpWorker) Run() {
	for {
		select {
		case <-w.cancel:
			return
		default:
			err := w.client.Send()
			if err != nil {

			} else {
				atomic.AddInt32(&w.monitor.Count, 1)
			}
		}

	}
}

func (w *HttpWorker) CheckInput(input interface{}) bool {
	if _, ok := input.(HttpInput); ok {
		return true
	}
	return false
}
