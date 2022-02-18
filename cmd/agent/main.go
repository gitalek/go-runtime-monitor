package main

import (
	"log"

	agent "github.com/gitalek/go-runtime-monitor/cmd/agent/app"
	"github.com/gitalek/go-runtime-monitor/internal/config"
)

func main() {
	cfg, err := config.Load(config.Path)
	if err != nil {
		log.Fatalln(err)
	}

	app := agent.NewApplication(cfg)
	log.Println("running the agent...")
	app.Run()
}
