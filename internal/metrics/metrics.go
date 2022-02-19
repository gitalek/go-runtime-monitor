package metrics

type IMetric interface {
	Kind() string
	Name() string
	StringifyValue() string
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

func (m Metric) Kind() string {
	return string(m.kind)
}
