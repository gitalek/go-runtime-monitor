package metrics

func NewMetricRandomValue(value float64) Gauge { return NewGauge(MetricRandomValue, value) }

func NewMetricPollCount(value int64) Counter { return NewCounter(MetricPollCount, value) }

func NewMetricAlloc(value float64) Gauge         { return NewGauge(MetricAlloc, value) }
func NewMetricBuckHashSys(value float64) Gauge   { return NewGauge(MetricBuckHashSys, value) }
func NewMetricFrees(value float64) Gauge         { return NewGauge(MetricFrees, value) }
func NewMetricGCCPUFraction(value float64) Gauge { return NewGauge(MetricGCCPUFraction, value) }
func NewMetricGCSys(value float64) Gauge         { return NewGauge(MetricGCSys, value) }
func NewMetricHeapAlloc(value float64) Gauge     { return NewGauge(MetricHeapAlloc, value) }
func NewMetricHeapIdle(value float64) Gauge      { return NewGauge(MetricHeapIdle, value) }
func NewMetricHeapInuse(value float64) Gauge     { return NewGauge(MetricHeapInuse, value) }
func NewMetricHeapObjects(value float64) Gauge   { return NewGauge(MetricHeapObjects, value) }
func NewMetricHeapReleased(value float64) Gauge  { return NewGauge(MetricHeapReleased, value) }
func NewMetricHeapSys(value float64) Gauge       { return NewGauge(MetricHeapSys, value) }
func NewMetricLastGC(value float64) Gauge        { return NewGauge(MetricLastGC, value) }
func NewMetricLookups(value float64) Gauge       { return NewGauge(MetricLookups, value) }
func NewMetricMCacheInuse(value float64) Gauge   { return NewGauge(MetricMCacheInuse, value) }
func NewMetricMCacheSys(value float64) Gauge     { return NewGauge(MetricMCacheSys, value) }
func NewMetricMSpanInuse(value float64) Gauge    { return NewGauge(MetricMSpanInuse, value) }
func NewMetricMSpanSys(value float64) Gauge      { return NewGauge(MetricMSpanSys, value) }
func NewMetricMallocs(value float64) Gauge       { return NewGauge(MetricMallocs, value) }
func NewMetricNextGC(value float64) Gauge        { return NewGauge(MetricNextGC, value) }
func NewMetricNumForcedGC(value float64) Gauge   { return NewGauge(MetricNumForcedGC, value) }
func NewMetricNumGC(value float64) Gauge         { return NewGauge(MetricNumGC, value) }
func NewMetricOtherSys(value float64) Gauge      { return NewGauge(MetricOtherSys, value) }
func NewMetricPauseTotalNs(value float64) Gauge  { return NewGauge(MetricPauseTotalNs, value) }
func NewMetricStackInuse(value float64) Gauge    { return NewGauge(MetricStackInuse, value) }
func NewMetricStackSys(value float64) Gauge      { return NewGauge(MetricStackSys, value) }
func NewMetricSys(value float64) Gauge           { return NewGauge(MetricSys, value) }
func NewMetricTotalAlloc(value float64) Gauge    { return NewGauge(MetricTotalAlloc, value) }
