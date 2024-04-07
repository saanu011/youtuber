package worker

import (
	"fmt"
	"time"
	"youtuber/pkg/redis"

	"github.com/hibiken/asynq"
)

type Worker struct {
	conf      Config
	scheduler *asynq.Scheduler
}

func New(conf Config, redis redis.Config) *Worker {
	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		fmt.Println("Unable to load load correct timezone for scheduler, defaulting to UTC %v", err)
	}

	scheduler := asynq.NewScheduler(
		asynq.RedisClientOpt{Addr: conf.Redis.Addr},
		&asynq.SchedulerOpts{
			Location: loc,
		},
	)

	w := &Worker{
		conf:      conf,
		scheduler: scheduler,
	}

	return w
}

func (s *Worker) Start() error {

	go func() {
		fmt.Println("Starting scheduler")

		err := s.scheduler.Run()
		if err != nil {
			fmt.Println("Failed to start scheduler %v", err)
		}
	}()

	return nil
}
