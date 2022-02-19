package metrics

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gitalek/go-runtime-monitor/internal/metrics"
	errorsApp "github.com/gitalek/go-runtime-monitor/internal/server/errors"
	"github.com/go-chi/chi/v5"
)

// GetMetric processes the GET '/value/{kind}/{name}' route.
func (m Metrics) GetMetric(w http.ResponseWriter, r *http.Request) {
	k, name := chi.URLParam(r, "kind"), chi.URLParam(r, "name")
	kind := metrics.Kind(k)
	if ok := metrics.ListKinds.Exists(kind); !ok {
		http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
		return
	}

	switch kind {
	case metrics.KindGauge:
		m.getMetricGauge(w, r, name)
	case metrics.KindCounter:
		m.getMetricCounter(w, r, name)
	default:
		// logic error
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (m Metrics) getMetricGauge(w http.ResponseWriter, _ *http.Request, name string) {
	metric, err := m.gauges.Get(name)
	if err != nil {
		switch {
		case errors.Is(err, errorsApp.ErrMetricNotFound):
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	_, err = w.Write([]byte(fmt.Sprintf("%f", metric.Value())))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (m Metrics) getMetricCounter(w http.ResponseWriter, _ *http.Request, name string) {
	metric, err := m.counters.Get(name)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	_, err = w.Write([]byte(fmt.Sprintf("%d", metric.Value())))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
