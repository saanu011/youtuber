package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"youtuber/src/app"
	"youtuber/src/config"
)

func main() {
	conf, err := config.Load("application.sample.yml")
	if err != nil {
		log.Fatal("Error loading config file ", err)
	}

	newApp, _ := app.NewApp(conf)
	err = newApp.Start()
	if err != nil {
		log.Fatal("Error starting app", err)
	}

	sign := make(chan os.Signal, 1)

	signal.Notify(sign, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-sign

	err = newApp.Shutdown()
	if err != nil {
		log.Fatal("Error shutting down serving", err)
	}
}
