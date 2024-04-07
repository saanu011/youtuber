package worker

import (
	"fmt"

	"github.com/hibiken/asynq"
)

const everyTenthMinuteSchedule = "*/10 * * * *"

func (w *Worker) RegisterTask(task *asynq.Task) {
	_, err := w.scheduler.Register(everyTenthMinuteSchedule, task)
	if err != nil {
		fmt.Println("failed to register task to scheduler: ", err)
	}
}
