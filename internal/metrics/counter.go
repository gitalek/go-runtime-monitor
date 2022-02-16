package metrics

import "fmt"

type Counter struct {
	Metric
	value int64
}

var _ IMetric = Counter{}

func NewCounter(name string, value int64) Counter {
	return Counter{Metric: NewMetric(name, KindCounter), value: value}
}

func (r Counter) Value() string {
	return fmt.Sprintf("%d", r.value)
}

func (r Counter) String() string {
	return fmt.Sprintf("type: %s, name: %s, value: %d", r.Type(), r.Name(), r.value)
}
