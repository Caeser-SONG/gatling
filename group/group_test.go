package group

import (
	"gatling/worker"
	"testing"
)

func TestGroup(t *testing.T) {
	input := &worker.HttpInput{
		Method: "GET",
		Url:    "http://www.baidu.com",
		Header: nil,
		Data:   nil,
	}
	var ch chan struct{}
	g := NewGroup("http", 100, input)
	g.StartWork()

	<-ch
}
