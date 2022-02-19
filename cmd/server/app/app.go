package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gitalek/go-runtime-monitor/internal/config"
)

type Application struct {
	cfg config.Config
}

func NewApplication(cfg config.Config) Application {
	return Application{cfg: cfg}
}

func (app Application) Run() {
	addr := app.cfg.Server.Host + ":" + app.cfg.Server.Port
	server := &http.Server{
		Addr:         addr,
		ReadTimeout:  time.Duration(app.cfg.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(app.cfg.Server.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(app.cfg.Server.IdleTimeout) * time.Second,
		Handler:      app.Routes(),
	}

	idleConnsClosed := make(chan struct{})

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
		sig := <-quit
		log.Printf("\ncaught shutdown signal: %s\n", sig.String())

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := server.Shutdown(ctx)
		if err != nil {
			log.Printf("server shutdown error: %s\n", err)
		}

		close(idleConnsClosed)
	}()

	log.Printf("START server on: %s\n", addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("server error: %s\n", err)
	}

	<-idleConnsClosed
	log.Printf("server has been STOPPED on: : %s\n", addr)
}
