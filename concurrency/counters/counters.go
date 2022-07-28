package counters

import (
	"sync"
	"sync/atomic"
)

type counterWithMutex struct {
	mu    sync.Mutex
	value int64
}

type counterWithRWMutex struct {
	rwMu  sync.RWMutex
	value int64
}

type counterWithAtomic struct {
	value int64
}

func (c *counterWithMutex) read() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func (c *counterWithRWMutex) inc() {
	c.rwMu.Lock()
	defer c.rwMu.Unlock()
	c.value++
}

func (c *counterWithRWMutex) read() int64 {
	c.rwMu.RLock()
	defer c.rwMu.RUnlock()
	return c.value
}

func (c *counterWithMutex) inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *counterWithAtomic) read() int64 {
	return c.value
}

func (c *counterWithAtomic) inc() {
	atomic.AddInt64(&c.value, 1)
}

type counters interface {
	read() int64
	inc()
}

func changeCounters(counters counters, numbers int) {

	for i := 0; i < numbers; i++ {
		go func() {
			for i := 0; i < numbers; i++ {
				counters.inc()
			}
		}()
	}
	for i := 0; i < numbers; i++ {
		go func() {
			for i := 0; i < numbers; i++ {
				counters.read()
			}
		}()

	}
}
