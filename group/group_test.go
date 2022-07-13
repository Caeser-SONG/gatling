package group

import (
	"gatling/tool"
	"gatling/worker"
	"testing"
)

func TestGroup(t *testing.T) {
	l := tool.NewLimiter(10)
	input := &worker.HttpInput{
		Method: "GET",
		Url:    "http://www.baidu.com",
		Header: nil,
		Data:   nil,
	}
	var ch chan struct{}
	g := NewGroup("http", 10, input, l)
	g.StartWork()

	<-ch
}
