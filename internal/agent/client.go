package agent

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/gitalek/go-runtime-monitor/internal/config"
	"github.com/gitalek/go-runtime-monitor/internal/metrics"
)

type Client struct {
	client http.Client
	cfg    config.Config
}

func NewClient(cfg config.Config) Client {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.MaxIdleConns = 100
	transport.MaxConnsPerHost = 100
	transport.MaxIdleConnsPerHost = 100
	client := http.Client{
		Timeout:   time.Duration(cfg.Agent.Timeout) * time.Second,
		Transport: transport,
	}
	return Client{client: client, cfg: cfg}
}

func (c Client) ReportMetric(wg *sync.WaitGroup, metric metrics.IMetric) error {
	defer wg.Done()

	url := c.generateURL(metric)
	request, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "text/plain")

	resp, err := c.client.Do(request)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	_, err = io.Copy(io.Discard, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func (c Client) Close() {
	c.client.CloseIdleConnections()
}

func (c Client) generateURL(metric metrics.IMetric) string {
	url := fmt.Sprintf(
		"http://%s:%s/update/%s/%s/%s",
		c.cfg.Server.Host,
		c.cfg.Server.Port,
		metric.Type(),
		metric.Name(),
		metric.Value(),
	)
	return url
}
