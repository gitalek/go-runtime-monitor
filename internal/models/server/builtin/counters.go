package builtin

import (
	"sync"

	"github.com/gitalek/go-runtime-monitor/internal/metrics"
	"github.com/gitalek/go-runtime-monitor/internal/models/server"
)

type StorageCounters struct {
	mu      sync.Mutex
	storage map[string]metrics.Counter
}

var _ server.IStorageCounters = &StorageCounters{}

func NewStorageCounters() *StorageCounters {
	return &StorageCounters{storage: make(map[string]metrics.Counter)}
}

func (s *StorageCounters) Update(metric metrics.Counter) {
	s.mu.Lock()
	defer s.mu.Unlock()
	m, ok := s.storage[metric.Name()]
	if !ok {
		s.storage[metric.Name()] = metric
		return
	}
	updated := metrics.NewCounter(m.Name(), m.Value()+metric.Value())
	s.storage[metric.Name()] = updated
}

func (s *StorageCounters) GetAll() map[string]metrics.Counter {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.storage
}
