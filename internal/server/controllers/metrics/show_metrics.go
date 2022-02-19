package metrics

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// ShowMetrics processes the GET '/' route.
func (m Metrics) ShowMetrics(w http.ResponseWriter, r *http.Request) {
	gauges := m.gauges.GetAll()
	counters := m.counters.GetAll()
	data := &templateData{
		Gauges:   gauges,
		Counters: counters,
	}

	path := filepath.Join("internal", "server", "views", "html", "metrics.page.tmpl")
	ts, err := template.ParseFiles(path)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	buf := new(bytes.Buffer)
	err = ts.ExecuteTemplate(buf, "metrics", data)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
