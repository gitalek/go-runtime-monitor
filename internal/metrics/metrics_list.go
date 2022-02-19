package metrics

const (
	MetricRandomValue = "RandomValue"

	MetricPollCount = "PollCount"

	MetricAlloc         = "Alloc"
	MetricBuckHashSys   = "BuckHashSys"
	MetricFrees         = "Frees"
	MetricGCCPUFraction = "GCCPUFraction"
	MetricGCSys         = "GCSys"
	MetricHeapAlloc     = "HeapAlloc"
	MetricHeapIdle      = "HeapIdle"
	MetricHeapInuse     = "HeapInuse"
	MetricHeapObjects   = "HeapObjects"
	MetricHeapReleased  = "HeapReleased"
	MetricHeapSys       = "HeapSys"
	MetricLastGC        = "LastGC"
	MetricLookups       = "Lookups"
	MetricMCacheInuse   = "MCacheInuse"
	MetricMCacheSys     = "MCacheSys"
	MetricMSpanInuse    = "MSpanInuse"
	MetricMSpanSys      = "MSpanSys"
	MetricMallocs       = "Mallocs"
	MetricNextGC        = "NextGC"
	MetricNumForcedGC   = "NumForcedGC"
	MetricNumGC         = "NumGC"
	MetricOtherSys      = "OtherSys"
	MetricPauseTotalNs  = "PauseTotalNs"
	MetricStackInuse    = "StackInuse"
	MetricStackSys      = "StackSys"
	MetricSys           = "Sys"
	MetricTotalAlloc    = "TotalAlloc"
)

type ListMetrics map[string]Kind

func (l ListMetrics) Exists(metric string) bool {
	_, ok := l[metric]
	return ok
}

func (l ListMetrics) Kind(metric string) (Kind, bool) {
	k, ok := l[metric]
	return k, ok
}

func (l ListMetrics) Count() int {
	return len(l)
}

var MetricKindList = ListMetrics{
	MetricRandomValue:   KindGauge,
	MetricPollCount:     KindCounter,
	MetricAlloc:         KindGauge,
	MetricBuckHashSys:   KindGauge,
	MetricFrees:         KindGauge,
	MetricGCCPUFraction: KindGauge,
	MetricGCSys:         KindGauge,
	MetricHeapAlloc:     KindGauge,
	MetricHeapIdle:      KindGauge,
	MetricHeapInuse:     KindGauge,
	MetricHeapObjects:   KindGauge,
	MetricHeapReleased:  KindGauge,
	MetricHeapSys:       KindGauge,
	MetricLastGC:        KindGauge,
	MetricLookups:       KindGauge,
	MetricMCacheInuse:   KindGauge,
	MetricMCacheSys:     KindGauge,
	MetricMSpanInuse:    KindGauge,
	MetricMSpanSys:      KindGauge,
	MetricMallocs:       KindGauge,
	MetricNextGC:        KindGauge,
	MetricNumForcedGC:   KindGauge,
	MetricNumGC:         KindGauge,
	MetricOtherSys:      KindGauge,
	MetricPauseTotalNs:  KindGauge,
	MetricStackInuse:    KindGauge,
	MetricStackSys:      KindGauge,
	MetricSys:           KindGauge,
	MetricTotalAlloc:    KindGauge,
}
