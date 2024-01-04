package main

import (
	"github.com/qwenode/gogo/ratelimit"
	"time"
)

func main() {
	limit := ratelimit.NewRateLimit(5)
	limit.Add()
	go func(l ratelimit.RateLimit) {
		defer l.Done()
		time.Sleep(time.Second * 5)
	}(limit)
	limit.Wait()
}
