package agent

import (
	"context"
	"runtime"
	"time"

	"github.com/gitalek/go-runtime-monitor/internal/metrics"
	"github.com/gitalek/go-runtime-monitor/internal/models"
)

type Poller struct {
	storage     models.IStorage
	pollCount   *PollCounter
	randomValue *RandomValueGenerator
}

func NewPoller(storage models.IStorage) Poller {
	return Poller{
		storage:     storage,
		pollCount:   NewPollCount(),
		randomValue: NewRandomValue(),
	}
}

func (p Poller) SchedulePoll(ctx context.Context, pollInterval int) {
	d := time.Duration(pollInterval) * time.Second
	ticker := time.NewTicker(d)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			data := p.poll()
			p.storage.SetMetrics(data)
		case <-ctx.Done():
			return
		}
	}
}
func (p Poller) poll() []metrics.IMetric {
	var rtm runtime.MemStats
	runtime.ReadMemStats(&rtm)

	result := make([]metrics.IMetric, 0, metrics.Count)

	result = append(result, p.randomValue.Generate())

	result = append(result, p.pollCount.Up())

	result = append(result, metrics.NewMetricAlloc(float64(rtm.Alloc)))
	result = append(result, metrics.NewMetricBuckHashSys(float64(rtm.BuckHashSys)))
	result = append(result, metrics.NewMetricFrees(float64(rtm.Frees)))
	result = append(result, metrics.NewMetricGCCPUFraction(rtm.GCCPUFraction))
	result = append(result, metrics.NewMetricGCSys(float64(rtm.GCSys)))
	result = append(result, metrics.NewMetricHeapAlloc(float64(rtm.HeapAlloc)))
	result = append(result, metrics.NewMetricHeapIdle(float64(rtm.HeapIdle)))
	result = append(result, metrics.NewMetricHeapInuse(float64(rtm.HeapInuse)))
	result = append(result, metrics.NewMetricHeapObjects(float64(rtm.HeapObjects)))
	result = append(result, metrics.NewMetricHeapReleased(float64(rtm.HeapReleased)))
	result = append(result, metrics.NewMetricHeapSys(float64(rtm.HeapSys)))
	result = append(result, metrics.NewMetricLastGC(float64(rtm.LastGC)))
	result = append(result, metrics.NewMetricLookups(float64(rtm.Lookups)))
	result = append(result, metrics.NewMetricMCacheInuse(float64(rtm.MCacheInuse)))
	result = append(result, metrics.NewMetricMCacheSys(float64(rtm.MCacheSys)))
	result = append(result, metrics.NewMetricMSpanInuse(float64(rtm.MSpanInuse)))
	result = append(result, metrics.NewMetricMSpanSys(float64(rtm.MSpanSys)))
	result = append(result, metrics.NewMetricMallocs(float64(rtm.Mallocs)))
	result = append(result, metrics.NewMetricNextGC(float64(rtm.NextGC)))
	result = append(result, metrics.NewMetricNumForcedGC(float64(rtm.NumForcedGC)))
	result = append(result, metrics.NewMetricNumGC(float64(rtm.NumGC)))
	result = append(result, metrics.NewMetricOtherSys(float64(rtm.OtherSys)))
	result = append(result, metrics.NewMetricPauseTotalNs(float64(rtm.PauseTotalNs)))
	result = append(result, metrics.NewMetricStackInuse(float64(rtm.StackInuse)))
	result = append(result, metrics.NewMetricStackSys(float64(rtm.StackSys)))
	result = append(result, metrics.NewMetricSys(float64(rtm.Sys)))
	result = append(result, metrics.NewMetricTotalAlloc(float64(rtm.TotalAlloc)))

	return result
}
