package metrics

import (
	"fmt"
)

type Gauge struct {
	Metric
	value float64
}

var _ IMetric = Gauge{}

func NewGauge(name string, value float64) Gauge {
	return Gauge{Metric: NewMetric(name, KindGauge), value: value}
}

func (r Gauge) Value() string {
	return fmt.Sprintf("%f", r.value)
}

func (r Gauge) String() string {
	return fmt.Sprintf("type: %s, name: %s, value: %f", r.Type(), r.Name(), r.value)
}
