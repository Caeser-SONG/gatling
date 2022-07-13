package tool

import (
	"golang.org/x/time/rate"
)

func NewLimiter(qps int64) *rate.Limiter {
	r := rate.Limit(qps)
	return rate.NewLimiter(r, 100)
}
