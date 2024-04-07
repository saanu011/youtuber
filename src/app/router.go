package app

import (
	"github.com/gorilla/mux"
	"github.com/hibiken/asynq"
	"net/http"

	"youtuber/src/config"
	"youtuber/src/resource/handler"
	"youtuber/src/resource/job"
)

func NewRouter(conf *config.Config, deps *Dependencies) http.Handler {
	router := mux.NewRouter()

	resourceHandler := handler.NewHandler(conf.Resource, deps.ResourceService)

	router.HandleFunc("/videos", resourceHandler.GetVideoListByParams).Methods(http.MethodGet)

	return router
}

func NewWorkerRouter(conf *config.Config, deps *Dependencies) asynq.Handler {

	router := asynq.NewServeMux()

	refreshDataJob := job.NewRefreshDataJob(deps.ResourceService)
	router.HandleFunc(job.RefreshDataTask, refreshDataJob.HandleRefreshDataTask)

	return router
}
