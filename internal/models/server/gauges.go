package server

import "github.com/gitalek/go-runtime-monitor/internal/metrics"

type IStorageGauges interface {
	Set(metric metrics.Gauge)
	Get(name string) (metrics.Gauge, error)
	GetAll() []metrics.Gauge
}
