package config

import (
	"github.com/spf13/viper"
	"youtuber/pkg/redis"
	"youtuber/pkg/server"
	"youtuber/pkg/worker"
	"youtuber/src/external/youtube"
	"youtuber/src/resource/handler"
)

type Config struct {
	Youtube  youtube.Config
	Resource handler.Config
	Server   server.Config
	Worker   worker.Config
	Redis    redis.Config
}

func Load(configFile string) (*Config, error) {
	v := viper.New()
	v.AutomaticEnv()

	v.SetConfigName(configFile)
	v.SetConfigType("yml")

	configPath := v.GetString("CONFIG_PATH")
	if configPath != "" {
		v.AddConfigPath(configPath)
	}

	v.AddConfigPath("config")

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config Config
	unmarshalErr := v.Unmarshal(&config)

	return &config, unmarshalErr
}
