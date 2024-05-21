package main

import "sync/atomic"

type AtomicCounter struct {
	cnt int64
}

func NewAtomicCounter() *AtomicCounter {
	return &AtomicCounter{
		cnt: 0,
	}
}

func (c *AtomicCounter) Get() interface{} {
	return atomic.LoadInt64(&c.cnt)
}

func (c *AtomicCounter) Inc() {
	atomic.AddInt64(&c.cnt, 1)
}
