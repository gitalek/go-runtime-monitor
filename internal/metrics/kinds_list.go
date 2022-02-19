package metrics

const (
	KindGauge   Kind = "gauge"
	KindCounter Kind = "counter"
)

type Kind string

type ListKinds map[Kind]struct{}

func (l ListKinds) Exists(kind Kind) bool {
	_, ok := l[kind]
	return ok
}

func (l ListKinds) Count() int {
	return len(l)
}
