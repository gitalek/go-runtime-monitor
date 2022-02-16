package agent

import (
	"math/rand"
	"sync"
	"time"

	"github.com/gitalek/go-runtime-monitor/internal/metrics"
)

type RandomValueGenerator struct {
	mu    sync.Mutex
	value float64
}

func NewRandomValue() *RandomValueGenerator {
	rand.Seed(time.Now().UnixNano())
	return &RandomValueGenerator{value: rand.Float64()} // #nosec G404 -- no need for cryptographic security here
}

func (g *RandomValueGenerator) Generate() metrics.Gauge {
	g.mu.Lock()
	g.value = rand.Float64() // #nosec G404 -- no need for cryptographic security here
	value := g.value
	g.mu.Unlock()
	return metrics.NewMetricRandomValue(value)
}

func (g *RandomValueGenerator) CurrentValue() metrics.Gauge {
	g.mu.Lock()
	value := g.value
	g.mu.Unlock()
	return metrics.NewMetricRandomValue(value)
}
