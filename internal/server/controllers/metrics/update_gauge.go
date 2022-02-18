package metrics

import (
	"net/http"
	"strconv"

	"github.com/gitalek/go-runtime-monitor/internal/metrics"
	"github.com/go-chi/chi/v5"
)

// UpdateGauge processes the POST '/update/gauge/{name}/{value}' route.
func (m Metrics) UpdateGauge(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	name, value := chi.URLParam(r, "name"), chi.URLParam(r, "value")
	v, err := normalizeGaugeInput(value)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	metric := metrics.NewGauge(name, v)
	m.gauges.Set(metric)
	w.WriteHeader(http.StatusOK)
}

func normalizeGaugeInput(value string) (float64, error) {
	return strconv.ParseFloat(value, 64)
}
