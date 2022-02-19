package builtin

import (
	"sync"

	"github.com/gitalek/go-runtime-monitor/internal/metrics"
	"github.com/gitalek/go-runtime-monitor/internal/models/server"
	"github.com/gitalek/go-runtime-monitor/internal/server/errors"
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

func (s *StorageCounters) Get(name string) (metrics.Counter, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	m, ok := s.storage[name]
	if !ok {
		return metrics.Counter{}, errors.ErrMetricNotFound
	}
	return m, nil
}

func (s *StorageCounters) GetAll() []metrics.Counter {
	s.mu.Lock()
	defer s.mu.Unlock()
	counters := make([]metrics.Counter, 0, len(s.storage))
	for _, c := range s.storage {
		counters = append(counters, c)
	}
	return counters
}
