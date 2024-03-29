package group

import (
	"fmt"
	"gatling/monitor"
	"gatling/worker"

	"golang.org/x/time/rate"
)

const (
	start = 100
)

type Group struct {
	Workers []worker.Worker
	Monitor *monitor.Monitor
	Qps     int
	proto   string
}

func NewGroup(proto string, Qps int, input *worker.HttpInput, limit *rate.Limiter) *Group {

	if !worker.CheckHttpInput(input) {
		fmt.Println("input is wrong")
		return nil
	}
	workers := make([]worker.Worker, 0)

	m := monitor.NewMonitor(input.Url)

	if proto == "http" {

		for i := 0; i < start; i++ {
			newworker := worker.NewHttpWorker(input, m, limit)
			workers = append(workers, newworker)
		}
	}

	return &Group{
		Workers: workers,
		Qps:     Qps,
		Monitor: m,
	}
}

func (g *Group) StartWork() {
	go g.Monitor.Watch()
	fmt.Println(len(g.Workers))
	for _, worker := range g.Workers {
		go worker.Run()
	}
}
