package metrics_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	tp "github.com/gitalek/go-runtime-monitor/internal/testing"
	"github.com/stretchr/testify/assert"
)

func testUpdateGauge(ts *httptest.Server) func(t *testing.T) {
	path := tp.GenerateUpdateGaugePath("TotalAlloc", "3556944")
	return func(t *testing.T) {
		type want struct {
			code        int
			contentType string
			body        string
		}
		tests := []struct {
			name string
			want want
		}{
			{
				name: "#1",
				want: want{
					code:        http.StatusOK,
					contentType: "text/plain; charset=utf-8",
					body:        "OK",
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				resp, body := tp.Request(t, ts, http.MethodPost, path)
				defer resp.Body.Close()
				assert.Equal(t, tt.want.code, resp.StatusCode)
				assert.Equal(t, tt.want.contentType, resp.Header.Get("Content-Type"))
				assert.Equal(t, tt.want.body, tp.TrimRespBodyString(body))
			})
		}
	}
}

func testUpdateGaugeFailedValidation(ts *httptest.Server) func(t *testing.T) {
	return func(t *testing.T) {
		type urlParams struct {
			name  string
			value string
		}
		type want struct {
			code        int
			contentType string
			body        string
		}
		w := want{
			code:        http.StatusBadRequest,
			contentType: "text/plain; charset=utf-8",
			body:        "Bad Request",
		}
		tests := []struct {
			name      string
			urlParams urlParams
			want      want
		}{
			{
				name: "non-float value",
				urlParams: urlParams{
					name:  "TotalAlloc",
					value: "two.one",
				},
				want: w,
			},
			{
				name: "float value (comma)",
				urlParams: urlParams{
					name:  "TotalAlloc",
					value: "2,7",
				},
				want: w,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				path := tp.GenerateUpdateGaugePath(tt.urlParams.name, tt.urlParams.value)
				resp, body := tp.Request(t, ts, http.MethodPost, path)
				defer resp.Body.Close()
				assert.Equal(t, tt.want.code, resp.StatusCode)
				assert.Equal(t, tt.want.contentType, resp.Header.Get("Content-Type"))
				assert.Equal(t, tt.want.body, tp.TrimRespBodyString(body))
			})
		}
	}
}
