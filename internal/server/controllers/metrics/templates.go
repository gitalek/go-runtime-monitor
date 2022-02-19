package metrics

import "github.com/gitalek/go-runtime-monitor/internal/metrics"

type templateData struct {
	Gauges   []metrics.Gauge
	Counters []metrics.Counter
}
