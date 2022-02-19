package metrics

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gitalek/go-runtime-monitor/internal/metrics"
	"github.com/go-chi/chi/v5"
)

// UpdateCounter processes the POST '/update/counter/{name}/{value}' route.
func (m Metrics) UpdateCounter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	name, value := chi.URLParam(r, "name"), chi.URLParam(r, "value")
	v, err := validateCounterInput(value)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	metric := metrics.NewCounter(name, v)
	m.counters.Update(metric)

	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprintln(w, http.StatusText(http.StatusOK))
	if err != nil {
		log.Println(err)
	}
}

func validateCounterInput(value string) (int64, error) {
	return strconv.ParseInt(value, 10, 64)
}
