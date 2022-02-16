package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gitalek/go-runtime-monitor/internal/agent"
	"github.com/gitalek/go-runtime-monitor/internal/config"
	"github.com/gitalek/go-runtime-monitor/internal/models/builtin"
)

type Application struct {
	cfg      config.Config
	poller   agent.Poller
	reporter agent.Reporter
}

func NewApplication(cfg config.Config) Application {
	storage := builtin.NewStorage()
	poller := agent.NewPoller(storage)
	client := agent.NewClient(cfg)
	reporter := agent.NewReporter(client, storage)
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

	go app.poller.SchedulePoll(ctx, app.cfg.PollInterval)
	go app.reporter.ScheduleReport(ctx, app.cfg.ReportInterval)

	<-ctx.Done()
}
