package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gitalek/go-runtime-monitor/internal/agent"
	"github.com/gitalek/go-runtime-monitor/internal/config"
	"github.com/gitalek/go-runtime-monitor/internal/models/agent/builtin"
)

type Application struct {
	cfg      config.Config
	poller   agent.Poller
	reporter agent.Reporter
}

func NewApplication(cfg config.Config) Application {
	storage := builtin.NewStorage()
	pollCount := agent.NewPollCount()
	client := agent.NewClient(cfg)

	var muPollerReporter sync.Mutex
	poller := agent.NewPoller(storage, pollCount, &muPollerReporter)
	reporter := agent.NewReporter(storage, pollCount, client, &muPollerReporter)
	return Application{
		cfg:      cfg,
		poller:   poller,
		reporter: reporter,
	}
}

func (app Application) Run() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
		sig := <-quit

		log.Printf("caught shutdown signal: %s\n", sig.String())
		cancel()
	}()

	go app.poller.SchedulePoll(ctx, app.cfg.Agent.PollInterval)
	go app.reporter.ScheduleReport(ctx, app.cfg.Agent.ReportInterval)

	<-ctx.Done()
}
