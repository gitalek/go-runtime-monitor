package agent

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/gitalek/go-runtime-monitor/internal/metrics"
	"github.com/gitalek/go-runtime-monitor/internal/models/agent"
)

type Reporter struct {
	// TODO: bottleneck
	// need for sync state (pollCount) between Reporter and Poller
	mu *sync.Mutex

	storage   agent.IStorage
	pollCount *PollCounter
	client    Client
}

func NewReporter(storage agent.IStorage, pollCount *PollCounter, client Client, mu *sync.Mutex) Reporter {
	return Reporter{
		mu:        mu,
		storage:   storage,
		pollCount: pollCount,
		client:    client,
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
			r.mu.Lock()
			data := r.storage.GetMetrics()
			r.pollCount.Reset()
			r.mu.Unlock()

			r.report(data)
		case <-ctx.Done():
			return
		}
	}
}

func (r Reporter) report(data []metrics.IMetric) {
	// TODO: need wait group?
	var wg sync.WaitGroup
	wg.Add(len(data))
	for _, metric := range data {
		go func(m metrics.IMetric) {
			err := r.client.ReportMetric(&wg, m)
			if err != nil {
				log.Println(err)
			}
		}(metric)
	}
	wg.Wait()
}
