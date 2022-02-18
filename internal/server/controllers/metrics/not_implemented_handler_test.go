package metrics_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	tp "github.com/gitalek/go-runtime-monitor/internal/testing"
	"github.com/stretchr/testify/assert"
)

func testNotImplementedHandler(ts *httptest.Server) func(t *testing.T) {
	return func(t *testing.T) {
		type urlParams struct {
			kind  string
			name  string
			value string
		}
		type want struct {
			code        int
			contentType string
			body        string
		}
		w := want{
			code:        http.StatusNotImplemented,
			contentType: "text/plain; charset=utf-8",
			body:        "Not Implemented",
		}
		tests := []struct {
			name      string
			urlParams urlParams
			want      want
		}{
			{
				name: "Unknown kind",
				urlParams: urlParams{
					kind:  "unknown",
					name:  "PollCount",
					value: "5079",
				},
				want: w,
			},
			{
				name: "Empty kind",
				urlParams: urlParams{
					kind:  "",
					name:  "PollCount",
					value: "5079",
				},
				want: w,
			},
			{
				name: "Whitespace kind",
				urlParams: urlParams{
					kind:  " ",
					name:  "PollCount",
					value: "5079",
				},
				want: w,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				path := tp.GenerateUpdatePath(tt.urlParams.kind, tt.urlParams.name, tt.urlParams.value)
				resp, body := tp.Request(t, ts, http.MethodPost, path)
				assert.Equal(t, tt.want.code, resp.StatusCode)
				assert.Equal(t, tt.want.contentType, resp.Header.Get("Content-Type"))
				assert.Equal(t, tt.want.body, tp.TrimRespBodyString(body))
			})
		}
	}
}
