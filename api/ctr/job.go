package ctr

import (
	"gogsWebHook/api/conf"
	"gogsWebHook/api/defs"
)

var (
	JobP *defs.JobProcess
)

func JobRun() {
	config := conf.Config
	JobP = &defs.JobProcess{
		JC:             make(chan string, 100),
		JavaPath:       config.JavaPath,
		JenkinsCliPath: config.JenkinsCliPath,
		JenkinsURL:     config.JenkinsURL,
		JenkinsAuth:    config.JenkinsAuth,
	}
	JobP.Run()
}
