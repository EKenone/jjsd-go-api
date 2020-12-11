package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"jjsd-go-api/api/mini/conf"
	"jjsd-go-api/api/mini/route"
	"log"
)

var env string

func init() {
	flag.StringVar(&env, "env", "debug", "default gin mode")
}

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		log.Panic(err)
	}
	gin.SetMode(env)

	eg := gin.Default()
	eg = route.Init(eg)
	eg.Run(":8081")
}
