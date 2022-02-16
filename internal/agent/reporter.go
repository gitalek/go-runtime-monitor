package agent

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/gitalek/go-runtime-monitor/internal/metrics"
	"github.com/gitalek/go-runtime-monitor/internal/models"
)

type Reporter struct {
	client  Client
	storage models.IStorage
}

func NewReporter(client Client, storage models.IStorage) Reporter {
	return Reporter{
		client:  client,
		storage: storage,
	}
}

func (r Reporter) ScheduleReport(ctx context.Context, reportInterval int) {
	d := time.Duration(reportInterval) * time.Second
	ticker := time.NewTicker(d)
	defer r.client.Close()
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			data := r.storage.GetMetrics()
			r.report(data)
		case <-ctx.Done():
			return
		}
	}
}

func (r Reporter) report(data []metrics.IMetric) {
	// ??
	var wg sync.WaitGroup
	wg.Add(len(data))
	for _, metric := range data {
		// ??
		go func(m metrics.IMetric) {
			err := r.client.ReportMetric(&wg, m)
			if err != nil {
				log.Println(err)
			}
		}(metric)
	}
	wg.Wait()
}
