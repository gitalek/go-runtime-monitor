package metrics

import (
	"github.com/gitalek/go-runtime-monitor/internal/models/server"
)

type Metrics struct {
	gauges   server.IStorageGauges
	counters server.IStorageCounters
}

func New(gauges server.IStorageGauges, counters server.IStorageCounters) Metrics {
	return Metrics{
		gauges:   gauges,
		counters: counters,
	}
}
