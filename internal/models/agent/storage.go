package agent

import (
	"github.com/gitalek/go-runtime-monitor/internal/metrics"
)

type IStorage interface {
	SetMetrics(m []metrics.IMetric)
	GetMetrics() []metrics.IMetric
}
