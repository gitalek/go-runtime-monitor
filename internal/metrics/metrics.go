package metrics

const (
	KindGauge   Kind = "gauge"
	KindCounter Kind = "counter"
)

const (
	Count = 29

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

type Kind string

type IMetric interface {
	Type() string
	Name() string
	Value() string
	String() string
}

type Metric struct {
	name string
	kind Kind
}

func NewMetric(name string, t Kind) Metric {
	return Metric{
		name: name,
		kind: t,
	}
}

func (m Metric) Name() string {
	return m.name
}

func (m Metric) Type() string {
	return string(m.kind)
}
