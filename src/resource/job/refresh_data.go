package job

import (
	"context"
	"fmt"
	"youtuber/pkg/redis"
	"youtuber/src/resource/domain"
	"youtuber/src/resource/service"

	"github.com/hibiken/asynq"
)

const RefreshDataTask = "refresh"

type RefreshDataJob struct {
	redisClient     redis.Client
	resourceService service.Service
}

var ResourceData []domain.Resource

func (j *RefreshDataJob) HandleRefreshDataTask(ctx context.Context, t *asynq.Task) error {
	var p domain.ResourceQuery = domain.ResourceQuery{
		ResourceType: "video",
		Query:        "football",
	}

	response, err := j.resourceService.GetResourcesList(ctx, p)
	if err != nil {
		fmt.Println("error fetching resource data", err)
	}

	ResourceData = response

	// TODO: add talk to database to upsert above fetched data

	return nil
}

func NewRefreshDataTask() (*asynq.Task, error) {
	return asynq.NewTask(RefreshDataTask, nil), nil
}

func NewRefreshDataJob(service service.Service) RefreshDataJob {

	return RefreshDataJob{
		resourceService: service,
	}
}
