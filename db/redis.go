package db

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"gopkg.in/ini.v1"
)

var (
	Redis *redis.Client
	Ctx   context.Context
)

func init() {
	file, err := ini.Load("./config/env.ini")
	if err != nil {
		fmt.Printf("env.ini读取失败：%v\n", err)
	}

	Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", file.Section("redis").Key("host"), file.Section("redis").Key("port")),
		Password: file.Section("redis").Key("password").String(),
		DB:       file.Section("redis").Key("database").MustInt(),
	})

	Ctx = context.Background()

	_, err = Redis.Ping(Ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("redis连接失败：%v", err))
	}
}
