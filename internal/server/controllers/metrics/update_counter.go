package metrics

import (
	"net/http"
	"strconv"

	"github.com/gitalek/go-runtime-monitor/internal/metrics"
	"github.com/go-chi/chi/v5"
)

// UpdateCounter processes the POST '/update/counter/{name}/{value}' route.
func (m Metrics) UpdateCounter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	name, value := chi.URLParam(r, "name"), chi.URLParam(r, "value")
	v, err := normalizeCounterInput(value)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	metric := metrics.NewCounter(name, v)
	m.counters.Update(metric)
	w.WriteHeader(http.StatusOK)
}

func normalizeCounterInput(value string) (int64, error) {
	return strconv.ParseInt(value, 10, 64)
}
