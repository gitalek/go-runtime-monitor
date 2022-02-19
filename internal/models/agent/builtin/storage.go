package builtin

import (
	"sync"

	"github.com/gitalek/go-runtime-monitor/internal/metrics"
	"github.com/gitalek/go-runtime-monitor/internal/models/agent"
)

type Metrics map[string]metrics.IMetric

type Storage struct {
	mu      sync.Mutex
	metrics Metrics
}

var _ agent.IStorage = &Storage{}

func NewStorage() *Storage {
	return &Storage{metrics: make(Metrics)}
}

func (s *Storage) SetMetrics(metrics []metrics.IMetric) {
	s.mu.Lock()
	defer s.mu.Unlock()
	target := Metrics{}
	for _, metric := range metrics {
		target[metric.Name()] = metric
	}
	s.metrics = target
}

func (s *Storage) GetMetrics() []metrics.IMetric {
	s.mu.Lock()
	defer s.mu.Unlock()
	result := make([]metrics.IMetric, 0, len(s.metrics))
	for _, metric := range s.metrics {
		if metric.Name() == metrics.MetricPollCount {
		}
		result = append(result, metric)
	}
	return result
}
