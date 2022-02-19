package builtin

import (
	"sync"

	"github.com/gitalek/go-runtime-monitor/internal/metrics"
	"github.com/gitalek/go-runtime-monitor/internal/models/server"
	"github.com/gitalek/go-runtime-monitor/internal/server/errors"
)

type StorageGauges struct {
	mu      sync.Mutex
	storage map[string]metrics.Gauge
}

var _ server.IStorageGauges = &StorageGauges{}

func NewStorageGauges() *StorageGauges {
	return &StorageGauges{storage: make(map[string]metrics.Gauge)}
}

func (s *StorageGauges) Set(metric metrics.Gauge) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.storage[metric.Name()] = metric
}

func (s *StorageGauges) Get(name string) (metrics.Gauge, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	m, ok := s.storage[name]
	if !ok {
		return metrics.Gauge{}, errors.ErrMetricNotFound
	}
	return m, nil
}

func (s *StorageGauges) GetAll() []metrics.Gauge {
	s.mu.Lock()
	defer s.mu.Unlock()
	gauges := make([]metrics.Gauge, 0, len(s.storage))
	for _, g := range s.storage {
		gauges = append(gauges, g)
	}
	return gauges
}
