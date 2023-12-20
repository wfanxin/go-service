package router

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wfanxin/go-service/db"
	"github.com/wfanxin/go-service/models"
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
		err := db.Redis.Set(db.Ctx, "name", "徐恩光", 0).Err()
		if err != nil {
			panic(err)
		}

		name, err := db.Redis.Get(db.Ctx, "name").Result()
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "redis" + name,
		})
	})

	r.GET("/mysql", func(c *gin.Context) {
		idstr := c.DefaultQuery("id", "0")
		id, _ := strconv.Atoi(idstr)
		user, _ := models.GetUser(id)

		fmt.Println(user)

		c.JSON(http.StatusOK, gin.H{
			"message": "mysql:id->" + idstr,
		})
	})

	return r
}
