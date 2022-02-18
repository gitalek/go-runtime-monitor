package app

import (
	"net/http"

	"github.com/gitalek/go-runtime-monitor/internal/models/server/builtin"
	"github.com/gitalek/go-runtime-monitor/internal/server/controllers/metrics"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app Application) Routes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.AllowContentType("text/plain"))

	gauges := builtin.NewStorageGauges()
	counters := builtin.NewStorageCounters()
	metricsController := metrics.New(gauges, counters)

	r.Post("/update/gauge/{name}/{value}", metricsController.UpdateGauge)
	r.Post("/update/counter/{name}/{value}", metricsController.UpdateCounter)
	r.Get("/show", metricsController.Show)

	return r
}
