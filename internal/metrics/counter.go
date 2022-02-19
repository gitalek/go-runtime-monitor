package metrics

import (
	"encoding/json"
	"fmt"
)

type Counter struct {
	Metric
	value int64
}

var _ IMetric = Counter{}

func NewCounter(name string, value int64) Counter {
	return Counter{Metric: NewMetric(name, KindCounter), value: value}
}

func (c Counter) StringifyValue() string {
	return fmt.Sprintf("%d", c.value)
}

func (c Counter) Value() int64 {
	return c.value
}

func (c Counter) String() string {
	return fmt.Sprintf("type: %s, name: %s, value: %d", c.Kind(), c.Name(), c.value)
}

func (c Counter) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Kind  Kind   `json:"kind"`
		Name  string `json:"name"`
		Value int64  `json:"value"`
	}{
		Kind:  c.kind,
		Name:  c.name,
		Value: c.value,
	})
}
