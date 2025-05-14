package redis

import (
	"context"

	"recsvc/internal/domain/er"

	"github.com/redis/go-redis/v9"
)

type RedisService struct {
	Cli *redis.Client
}

func New() (*RedisService, error) {
	cli := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if err := cli.Ping(context.TODO()).Err(); err != nil {
		return nil, er.W(err)
	}

	return &RedisService{
		Cli: cli,
	}, nil
}
