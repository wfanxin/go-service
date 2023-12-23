package main

import (
	"github.com/wfanxin/go-service/config"
	"github.com/wfanxin/go-service/router"
	"github.com/wfanxin/go-service/utils"
)

func main() {
	r := router.Router()

	utils.Log.Info("nihao")

	r.Run(config.Port)
}
