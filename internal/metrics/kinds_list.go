package metrics

const (
	KindGauge   Kind = "gauge"
	KindCounter Kind = "counter"
)

type Kind string

type ListK map[Kind]struct{}

func (l ListK) Exists(kind Kind) bool {
	_, ok := l[kind]
	return ok
}

func (l ListK) Count() int {
	return len(l)
}

var ListKinds = ListK{
	KindGauge:   struct{}{},
	KindCounter: struct{}{},
}
