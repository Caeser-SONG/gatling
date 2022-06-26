package group

import (
	"fmt"
	"gatling/worker"
	"runtime"
	"testing"
)

func TestGroup(t *testing.T) {
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(4)
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
