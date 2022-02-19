package metrics

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gitalek/go-runtime-monitor/internal/metrics"
	"github.com/go-chi/chi/v5"
)

// UpdateGauge processes the POST '/update/gauge/{name}/{value}' route.
func (m Metrics) UpdateGauge(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	name, value := chi.URLParam(r, "name"), chi.URLParam(r, "value")
	v, err := normalizeGaugeInput(value)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	metric := metrics.NewGauge(name, v)
	m.gauges.Set(metric)

	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprintln(w, http.StatusText(http.StatusOK))
	if err != nil {
		log.Println(err)
	}
}

func normalizeGaugeInput(value string) (float64, error) {
	n, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return n, err
	}
	roundedString := fmt.Sprintf("%.3f", n)
	n, err = strconv.ParseFloat(roundedString, 64)
	if err != nil {
		return n, err
	}
	return n, err
}
