package worker
0;95;0c
import (
	"gatling/client"
)

type Worker interface {
	Run()
}

type HttpWorker struct {
	header map[string]string
	url    string
	data   map[string]interface{}
	cancel chan struct{}
	client *client.HttpClient
}

func NewHttpWorker(method, url string, data map[string]interface{}, header map[string]string) *HttpWorker {

	
	return &HttpWorker{
		header: header,
		url:    url,
		data:   data,
		cancel: make(chan struct{}),
		client: client.NewHttpClient(method, url, data, header),
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
			w.client.Send()
		}

	}
}
