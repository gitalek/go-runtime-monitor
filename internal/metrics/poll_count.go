package metrics

import (
	"sync"
)

type PollCount struct {
	mu    sync.Mutex
	value int64
}

func NewPollCount() *PollCount {
	return &PollCount{value: 0}
}

func (c *PollCount) Up() Counter {
	c.mu.Lock()
	c.value++
	value := c.value
	c.mu.Unlock()
	return NewMetricPollCount(value)
}

func (c *PollCount) CurrentValue() Counter {
	c.mu.Lock()
	value := c.value
	c.mu.Unlock()
	return NewMetricPollCount(value)
}
