package main

import (
	"github.com/wfanxin/go-service/config"
	"github.com/wfanxin/go-service/router"
)

func main() {
	r := router.Router()

	r.Run(config.Port)
}
