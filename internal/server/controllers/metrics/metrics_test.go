package metrics_test

import (
	"net/http/httptest"
	"testing"

	tp "github.com/gitalek/go-runtime-monitor/internal/testing"
)

func setupWrapper(t *testing.T, name string, f func(ts *httptest.Server) func(t *testing.T)) {
	var teardown func()
	server, teardown := tp.NewServer()
	ts := httptest.NewServer(server.Routes())
	defer func() {
		teardown()
		ts.Close()
	}()

	t.Run(name, f(ts))
}

func TestMetrics_UpdateGauge(t *testing.T) {
	setupWrapper(t, "Should update gauge", testUpdateGauge)
	setupWrapper(t, "Failed validation", testUpdateGaugeFailedValidation)
}

func TestMetrics_UpdateCounter(t *testing.T) {
	setupWrapper(t, "Should update gauge", testUpdateCounter)
	setupWrapper(t, "Failed validation", testUpdateCounterFailedValidation)
}
