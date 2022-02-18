package metrics

import (
	"encoding/json"
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

func (g Gauge) StringifyValue() string {
	return fmt.Sprintf("%f", g.value)
}

func (g Gauge) Value() float64 {
	return g.value
}

func (g Gauge) String() string {
	return fmt.Sprintf("type: %s, name: %s, value: %f", g.Kind(), g.Name(), g.value)
}

func (g Gauge) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Kind  Kind    `json:"kind"`
		Name  string  `json:"name"`
		Value float64 `json:"value"`
	}{
		Kind:  g.kind,
		Name:  g.name,
		Value: g.value,
	})
}
