package group

import (
	"gatling/tool"
	"gatling/worker"
	"testing"
)

func aestLimitGroup(t *testing.T) {
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

func TestNoLimit(t *testing.T) {
	input := &worker.HttpInput{
		Method: "GET",
		Url:    "http://www.baidu.com",
	}

	var ch chan struct{}

	g := NewGroup("http", 100, input, nil)
	g.StartWork()
	<-ch
}
