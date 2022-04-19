package worker

import (
	"gatling/client"
	"net/http"
)

type Worker interface {
	Run()
}

type HttpWorker struct {
	header http.Header
	url    string
	proto  string
	client *client.HttpClient
}

func NewHttpWorker(url, header string) *HttpWorker {
	return &HttpWorker{}
}

// go  run this function
func (w *HttpWorker) Run() {
	for {
		w.client.Send()
	}
}
