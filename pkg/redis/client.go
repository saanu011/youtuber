package redis

import (
	"context"
	"fmt"

	redisV8 "github.com/go-redis/redis/v8"
)

type Client struct {
	*redisV8.Client
}

func NewClient(c Config) (Client, error) {
	redisClient := redisV8.NewClient(&redisV8.Options{
		Addr:     c.Address,
		Password: c.Password,
		Username: c.User,
	})

	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		fmt.Println("redis (address: %v) connection error: %v\n", c.Address, err)

		return Client{}, fmt.Errorf("redis connection error")
	}

	return Client{redisClient}, nil
}
