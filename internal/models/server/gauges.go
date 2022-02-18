package server

import "github.com/gitalek/go-runtime-monitor/internal/metrics"

type IStorageGauges interface {
	Set(metric metrics.Gauge)
	GetAll() map[string]metrics.Gauge
}
