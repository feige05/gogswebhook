package defs

import (
	"bytes"
	"fmt"
	"os/exec"
)

type JobProcess struct {
	JC             chan string
	JavaPath       string
	JenkinsCliPath string
	JenkinsURL     string
	JenkinsAuth    string
}

func (j *JobProcess) doExec(jobname string) {
	cmd := exec.Command(j.JavaPath, "-jar", j.JenkinsCliPath, "-s", j.JenkinsURL, "-auth", j.JenkinsAuth, "build", jobname)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	}
}

func (j *JobProcess) Run() {
	for jobname := range j.JC {
		fmt.Println("cmd::::", j.JavaPath, "-jar", j.JenkinsCliPath, "-s", j.JenkinsURL, "-auth", j.JenkinsAuth, "build", jobname)
		go j.doExec(jobname)
	}

}
