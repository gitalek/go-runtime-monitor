package main

import (
	"log"

	"github.com/gitalek/go-runtime-monitor/internal/config"
)

func main() {
	cfg, err := config.Load(config.Path)
	if err != nil {
		log.Fatalln(err)
	}

	app := NewApplication(cfg)
	log.Println("running the agent...")
	app.Run()
}
