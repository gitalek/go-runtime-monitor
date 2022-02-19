package metrics

import (
	"net/http"
)

// NotImplemented processes the POST '/update/{unknown_kind}/{name}/{value}' route.
func (m Metrics) NotImplemented(w http.ResponseWriter, _ *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}
