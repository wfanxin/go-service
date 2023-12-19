package router

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/wfanxin/go-service/config"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello",
		})
	})

	r.GET("/redis", func(c *gin.Context) {
		client := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%v:%v", config.Redis.Host, config.Redis.Port),
			Password: config.Redis.Password,
			DB:       config.Redis.Database,
		})
		ctx := context.Background()

		err := client.Set(ctx, "name", "徐恩光", 0).Err()
		if err != nil {
			panic(err)
		}

		name, err := client.Get(ctx, "name").Result()
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"message": name,
		})
	})

	return r
}
