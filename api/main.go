package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gogsWebHook/api/conf"
	"gogsWebHook/api/defs"
	"os"
	"gogsWebHook/api/ctr"
)

func main() {
	fmt.Println(os.Args[1:])

	var err error
	var config *defs.Config
	err = conf.Parse(os.Args[1:])
	config = conf.Config
	if err != nil || config == nil {
		fmt.Println("参数不正确！")
		os.Exit(0)
	}

	go ctr.JobRun()

	r := gin.Default()

	// Ping test
	r.GET("/ping", ctr.HandlePing)

	r.POST("/hook/:jobname", ctr.HandleJob)
	// Listen and Server in 0.0.0.0:8080
	r.Run(config.Port)
}
