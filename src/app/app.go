package app

import (
	"youtuber/pkg/server"
	"youtuber/pkg/worker"

	"youtuber/src/config"
)

type App struct {
	Server *server.Server
	Worker *worker.Worker
}

func NewApp(config *config.Config) (*App, error) {
	deps, err := NewDependencies(config)
	if err != nil {
		return nil, err
	}

	router := NewRouter(config, deps)

	httpServer := server.NewServer(config.Server, router)

	worker := worker.New(config.Worker, config.Redis)

	app := &App{
		Server: httpServer,
		Worker: worker,
	}

	return app, nil
}

func (a *App) Start() error {
	err := a.Server.Start()
	if err != nil {
		return err
	}

	return a.Worker.Start()
}

func (a *App) Shutdown() error {
	err := a.Server.Shutdown()

	return err
}
