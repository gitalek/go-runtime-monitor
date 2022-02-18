package server

import "github.com/gitalek/go-runtime-monitor/internal/metrics"

type IStorageCounters interface {
	Update(metric metrics.Counter)
	GetAll() map[string]metrics.Counter
}
