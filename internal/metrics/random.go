package metrics

import (
	"math/rand"
	"sync"
	"time"
)

type RandomValue struct {
	mu    sync.Mutex
	value float64
}

func NewRandomValue() *RandomValue {
	rand.Seed(time.Now().UnixNano())
	return &RandomValue{value: rand.Float64()}
}

func (r *RandomValue) Generate() Gauge {
	r.mu.Lock()
	r.value = rand.Float64()
	value := r.value
	r.mu.Unlock()
	return NewMetricRandomValue(value)
}

func (r *RandomValue) CurrentValue() Gauge {
	r.mu.Lock()
	value := r.value
	r.mu.Unlock()
	return NewMetricRandomValue(value)
}