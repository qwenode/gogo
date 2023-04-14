package ratelimit

import "sync"

type rateLimit struct {
    rate  chan int
    group *sync.WaitGroup
}

// NewRateLimit 速率限制功能
func NewRateLimit(limit int) *rateLimit {
    r := new(rateLimit)
    r.rate = make(chan int, limit)
    r.group = new(sync.WaitGroup)
    return r
}

// Add 添加一项
func (r *rateLimit) Add() {
    r.rate <- 1
    r.group.Add(1)
}

// Done 完成一项
func (r *rateLimit) Done() {
    r.group.Done()
    <-r.rate
}

// Wait 等待所有项目完成
func (r *rateLimit) Wait() {
    r.group.Wait()
}
