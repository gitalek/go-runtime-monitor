package agent

import (
	"sync"

	"github.com/gitalek/go-runtime-monitor/internal/metrics"
)

type PollCounter struct {
	mu    sync.Mutex
	value int64
}

func NewPollCount() *PollCounter {
	return &PollCounter{value: 0}
}

func (c *PollCounter) Up() metrics.Counter {
	c.mu.Lock()
	c.value++
	value := c.value
	c.mu.Unlock()
	return metrics.NewMetricPollCount(value)
}

func (c *PollCounter) Reset() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value = 0
}

func (c *PollCounter) CurrentValue() metrics.Counter {
	c.mu.Lock()
	value := c.value
	c.mu.Unlock()
	return metrics.NewMetricPollCount(value)
}
