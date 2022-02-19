package server

import "github.com/gitalek/go-runtime-monitor/internal/metrics"

type IStorageCounters interface {
	Update(metric metrics.Counter)
	Get(name string) (metrics.Counter, error)
	GetAll() []metrics.Counter
}
