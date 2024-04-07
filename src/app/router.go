package app

import (
	"github.com/gorilla/mux"
	"net/http"
	"youtuber/src/config"
	"youtuber/src/resource/handler"
)

func NewRouter(conf *config.Config, deps *Dependencies) http.Handler {
	router := mux.NewRouter()

	resourceHandler := handler.NewHandler(conf.Resource, deps.resourceService)

	router.HandleFunc("/videos", resourceHandler.GetVideoListByParams).Methods(http.MethodGet)

	return router
}
