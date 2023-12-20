package db

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/wfanxin/go-service/config"
)

var (
	Redis *redis.Client
	Ctx   context.Context
)

func init() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", config.Redis.Host, config.Redis.Port),
		Password: config.Redis.Password,
		DB:       config.Redis.Database,
	})

	Ctx = context.Background()
}
