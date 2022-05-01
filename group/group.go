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

func NewGroup(proto string, Qps int32, input worker.HttpInput) *Group {
	workers := make([]worker.Worker, 0)

	m := monitor.NewMonitor()
	if proto == "http" {
		for i := 0; i < start; i++ {
			newworker := worker.NewHttpWorker(input, m)
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
