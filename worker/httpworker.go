package worker

import (
	"fmt"
	"gatling/client"
	"gatling/monitor"
	"reflect"
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
				_, err := w.client.Send()
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
				_, err := w.client.Send()
				if err != nil {
					fmt.Println(err)
				}
				atomic.AddInt32(&w.monitor.Count, 1)
			}
		}
	}
}

func CheckHttpInput(input interface{}) bool {
	t := reflect.TypeOf(input)
	n := t.Name()
	fmt.Println(n)
	//	if n == "HttpInput" {
	//	return true
	//	}
	//	return false

	if _, ok := input.(*HttpInput); ok {
		return true
	}
	return false
}
