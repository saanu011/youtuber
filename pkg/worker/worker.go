package worker

import (
	"fmt"
	"time"

	"github.com/hibiken/asynq"
)

type Worker struct {
	conf      Config
	scheduler *asynq.Scheduler
	server    *asynq.Server
	handler   asynq.Handler
}

func New(conf Config, router asynq.Handler) *Worker {
	redisConnOpt := asynq.RedisClientOpt{Addr: conf.Redis.Addr}

	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		fmt.Println("Unable to load load correct timezone for scheduler, defaulting to UTC %v", err)
	}

	scheduler := asynq.NewScheduler(
		redisConnOpt,
		&asynq.SchedulerOpts{
			Location: loc,
		},
	)

	w := &Worker{
		conf:      conf,
		scheduler: scheduler,
		handler:   router,
	}

	server := asynq.NewServer(redisConnOpt, asynq.Config{
		Concurrency: conf.Concurrency,
	})

	w.server = server

	return w
}

func (w *Worker) Start() error {

	go func() {
		fmt.Println("Starting worker")

		err := w.server.Run(w.handler)
		if err != nil {
			fmt.Println("Failed to start worker", err)
		}
	}()

	go func() {
		fmt.Println("Starting scheduler")

		err := w.scheduler.Run()
		if err != nil {
			fmt.Println("Failed to start scheduler", err)
		}
	}()

	return nil
}

func (w *Worker) Shutdown() error {
	fmt.Println("Shutting down worker")
	w.server.Shutdown()

	return nil
}
