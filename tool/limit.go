package tool

import (
	"time"

	"golang.org/x/time/rate"
)

func NewLimiter(qps int) *rate.Limiter {
	t := time.Duration(qps)
	r := rate.Every(t)
	return rate.NewLimiter(r, qps*2)
}
