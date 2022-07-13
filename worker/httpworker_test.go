package worker

import (
	"gatling/monitor"
	"testing"

	"golang.org/x/time/rate"
)

func TestLimitworker(t *testing.T) {
	input := HttpInput{}

	NewHttpWorker(input*worker.HttpInput, monitor*monitor.Monitor, limiter*rate.Limiter)

}
