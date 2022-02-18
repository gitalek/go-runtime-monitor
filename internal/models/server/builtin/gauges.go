package builtin

import (
	"sync"

	"github.com/gitalek/go-runtime-monitor/internal/metrics"
	"github.com/gitalek/go-runtime-monitor/internal/models/server"
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

func (s *StorageGauges) GetAll() map[string]metrics.Gauge {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.storage
}
