package app

import (
	"youtuber/pkg/redis"
	"youtuber/src/config"
	"youtuber/src/external/youtube"
	"youtuber/src/resource/service"
)

type Dependencies struct {
	YoutubeAPIClient youtube.APIResponder
	ResourceService  service.Service
	RedisClient      redis.Client
}

func NewDependencies(conf *config.Config) (*Dependencies, error) {
	redisClient, err := redis.NewClient(conf.Redis)
	if err != nil {
		return nil, err
	}

	youtubeAPIClient := youtube.NewAPIClient(conf.Youtube)
	resourceService := service.NewService(youtubeAPIClient)

	return &Dependencies{
		YoutubeAPIClient: youtubeAPIClient,
		ResourceService:  resourceService,
		RedisClient:      redisClient,
	}, nil
}
