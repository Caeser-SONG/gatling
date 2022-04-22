package group

import (
	"gatling/monitor"
	"gatling/worker"
)

type Group struct {
	Workers []worker.Worker
	Monitor monitor.Monitor
	QPS     int32
}

func NewGroup() *Group {
	return &Group{}
}
