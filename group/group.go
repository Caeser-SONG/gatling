package group

import (
	"gatling/monitor"
	"gatling/worker"
)

const (
	start = 5
)

type Group struct {
	Workers []worker.Worker
	Monitor *monitor.Monitor
	Qps     int32
	proto   string
}

func NewGroup(method, url, proto string, Qps int32, header map[string]string, data map[string]interface{}) *Group {
	workers := make([]worker.Worker, 0)
	if proto == "http" {
		for i := 0; i < start; i++ {
			newworker := worker.NewHttpWorker(method, url, data, header)

			workers = append(workers, newworker)
		}
	}
	return &Group{
		Workers: workers,
		Qps:     Qps,
	}
}

func (g *Group) StartWork() {
	g.Monitor.Watch()
	for _, worker := range g.Workers {
		go worker.Run()
	}
}
