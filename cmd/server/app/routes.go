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

	r.Route("/update", func(r chi.Router) {
		r.Post("/gauge/{name}/{value}", metricsController.UpdateGauge)
		r.Post("/counter/{name}/{value}", metricsController.UpdateCounter)
		r.Post("/{unknown_kind}/{name}/{value}", metricsController.NotImplemented)
	})
	r.Get("/show", metricsController.Show)

	return r
}
