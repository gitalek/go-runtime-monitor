package builtin

import (
	"sync"

	"github.com/gitalek/go-runtime-monitor/internal/metrics"
	"github.com/gitalek/go-runtime-monitor/internal/models"
)

type Metrics map[string]metrics.IMetric

type Storage struct {
	mu      sync.Mutex
	metrics Metrics
}

var _ models.IStorage = &Storage{}

func NewStorage() *Storage {
	return &Storage{metrics: make(Metrics)}
}

func (s *Storage) SetMetrics(metrics []metrics.IMetric) {
	target := Metrics{}
	for _, metric := range metrics {
		target[metric.Name()] = metric
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.metrics = target
}

func (s *Storage) GetMetrics() []metrics.IMetric {
	s.mu.Lock()
	defer s.mu.Unlock()
	result := make([]metrics.IMetric, 0, len(s.metrics))
	for _, metric := range s.metrics {
		result = append(result, metric)
	}
	return result
}
