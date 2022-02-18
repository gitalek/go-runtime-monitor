package metrics

import (
	"encoding/json"
	"log"
	"net/http"
)

// Show processes the GET '/show' route.
func (m Metrics) Show(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	gauges := m.gauges.GetAll()
	counters := m.counters.GetAll()

	jsondataGauges, err := json.MarshalIndent(gauges, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsondataGauges = append(jsondataGauges, '\n')

	jsondataCounters, err := json.MarshalIndent(counters, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsondataCounters = append(jsondataCounters, '\n')

	jsondata := append(jsondataGauges, jsondataCounters...)

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsondata)
	if err != nil {
		log.Printf("error while writing json: %s", err)
	}
	// requestDump(r, "counter")
}
