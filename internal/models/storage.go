package models

import (
	"github.com/gitalek/go-runtime-monitor/internal/metrics"
)

type IStorage interface {
	SetMetrics(metrics []metrics.IMetric)
	GetMetrics() []metrics.IMetric
}
