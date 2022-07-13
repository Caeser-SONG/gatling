package worker

import (
	"fmt"
	"gatling/client"
	"gatling/monitor"
	"sync/atomic"
	"time"

	"golang.org/x/time/rate"
)

type HttpWorker struct {
	input   *HttpInput
	cancel  chan struct{}
	client  client.Client
	limiter *rate.Limiter
	monitor *monitor.Monitor
}

type HttpInput struct {
	Method string
	Url    string
	Header map[string]string
	Data   map[string]interface{}
}

func NewHttpWorker(input *HttpInput, monitor *monitor.Monitor, limiter *rate.Limiter) *HttpWorker {
	return &HttpWorker{
		input:   input,
		monitor: monitor,
		cancel:  make(chan struct{}),
		limiter: limiter,
		client:  client.NewHttpClient(input.Method, input.Url, input.Data, input.Header),
	}
}

// 停止一个并发
func (w *HttpWorker) Cancel() {
	w.cancel <- struct{}{}
}

// go  run this function
func (w *HttpWorker) Run() {
	if w.limiter == nil {

		for {

			select {
			case <-w.cancel:

				return
			default:
				start := time.Now()
				err := w.client.Send()
				cost := time.Since(start)

				if err != nil {
					fmt.Println(err)
					fmt.Println(cost)
				} else {
					// 计数
					atomic.AddInt32(&w.monitor.Count, 1)
					// 时间
				}
			}

		}
	} else {

		for {
			if w.limiter.Allow() {
				err := w.client.Send()
				if err != nil {
					fmt.Println(err)
				}
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
