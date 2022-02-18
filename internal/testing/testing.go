package testing

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"path"
	"runtime"
	"strings"
	"testing"

	server "github.com/gitalek/go-runtime-monitor/cmd/server/app"
	"github.com/gitalek/go-runtime-monitor/internal/config"
	"github.com/stretchr/testify/require"
)

const (
	configPath = "./config.yml"
)

func loadConfig() (config.Config, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("LoadConfig: unknown run filename")
	}

	dir := path.Join(path.Dir(filename), "../../")
	return config.Load(dir + strings.TrimLeft(configPath, "."))
}

func NewServer() (app server.Application, clean func()) {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	app = server.NewApplication(cfg)
	return app, func() {}
}

func Request(t *testing.T, ts *httptest.Server, method, urlPath string) (*http.Response, string) {
	req, err := http.NewRequest(method, ts.URL+urlPath, nil)
	require.NoError(t, err)

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)

	respBody, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)

	defer resp.Body.Close()

	return resp, string(respBody)
}
