package app

import (
	"fmt"
	"youtuber/pkg/server"
	"youtuber/pkg/worker"
	"youtuber/src/resource/job"

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

	fmt.Println(config.Resource.Host, config.Resource)

	serverRouter := NewRouter(config, deps)

	httpServer := server.New(config.Server, serverRouter)

	workerRouter := NewWorkerRouter(config, deps)
	work := worker.New(config.Worker, workerRouter)

	addSchedulerTask(work)

	app := &App{
		Server: httpServer,
		Worker: work,
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
	if err != nil {
		return err
	}
	return a.Worker.Shutdown()
}

func addSchedulerTask(work *worker.Worker) {
	task, _ := job.NewRefreshDataTask()
	work.RegisterTask(task)
}
