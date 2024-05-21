package main

import "sync"

type MutexCounter struct {
	mu  sync.Mutex
	cnt int
}

func NewMutexCounter() *MutexCounter {
	return &MutexCounter{
		cnt: 0,
	}
}

func (c *MutexCounter) Get() interface{} {
	return c.cnt
}

func (c *MutexCounter) Inc() {
	c.mu.Lock()
	c.cnt++
	c.mu.Unlock()
}
