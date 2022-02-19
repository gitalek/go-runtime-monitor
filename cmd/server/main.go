package main

import (
	"log"

	server "github.com/gitalek/go-runtime-monitor/cmd/server/app"
	"github.com/gitalek/go-runtime-monitor/internal/config"
)

func main() {
	cfg, err := config.Load(config.Path)
	if err != nil {
		log.Fatalln(err)
	}

	app := server.NewApplication(cfg)
	app.Run()
}
