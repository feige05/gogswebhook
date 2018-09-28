package ctr

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gogits/go-gogs-client"
	"github.com/pkg/errors"
	"gogsWebHook/api/conf"
	"net/http"
)

const (
	HOOK_EVENT_CREATE        string = "create"
	HOOK_EVENT_DELETE        string = "delete"
	HOOK_EVENT_FORK          string = "fork"
	HOOK_EVENT_PUSH          string = "push"
	HOOK_EVENT_ISSUES        string = "issues"
	HOOK_EVENT_PULL_REQUEST  string = "pull_request"
	HOOK_EVENT_ISSUE_COMMENT string = "issue_comment"
	HOOK_EVENT_RELEASE       string = "release"
)

func checkPayload(c *gin.Context, event, signature, secret string) error{
	var json gogs.Payloader
	switch event {
	case HOOK_EVENT_PUSH:
		json = &gogs.PushPayload{}
		break
	case HOOK_EVENT_RELEASE:
		json = &gogs.ReleasePayload{}
		break
	case HOOK_EVENT_CREATE:
		json = &gogs.CreatePayload{}
		break
	case HOOK_EVENT_DELETE:
		json = &gogs.DeletePayload{}
		break
	case HOOK_EVENT_FORK:
		json = &gogs.ForkPayload{}
		break
	case HOOK_EVENT_ISSUES:
		json = &gogs.IssuesPayload{}
		break
	case HOOK_EVENT_PULL_REQUEST:
		json = &gogs.PullRequestPayload{}
		break
	case HOOK_EVENT_ISSUE_COMMENT:
		json = &gogs.IssueCommentPayload{}
		break
	}
	if err := c.ShouldBindWith(json, binding.JSON); err != nil {
		return err
	}
	data,err := json.JSONPayload()

	if err != nil {
		return err
	}
	h := hmac.New(sha256.New, []byte(secret))
	if _, err := h.Write(data); err != nil {
		return err
	} else {
		sha := hex.EncodeToString(h.Sum(nil))
		if sha != signature {
			fmt.Println("gogs-signature: " + signature)
			fmt.Println("signature: " + sha)
			fmt.Println("gogs signature error")
			return  errors.New("gogs signature error")
		}
	}

	return nil
}


func getSecret(name string) string {
	cf := conf.Config
	for _, job := range cf.Store {
		if job.Name == name {
			return job.Secret
		}
	}
	return ""
}

func HandleJob(c *gin.Context) {
	jobname := c.Param("jobname")
	secret := getSecret(jobname)
	event := c.Request.Header.Get("X-Gogs-Event")
	signature := c.Request.Header.Get("X-Gogs-Signature")
	if secret == "" {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "job": jobname, "event": event, "signature": signature})
	} else if signature == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gogs signature error"})
	} else {
		if err := checkPayload(c, event, signature, secret); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "gogs signature error"})
		} else {
			JobP.JC <- jobname
			c.JSON(http.StatusOK, gin.H{"status": "ok", "job": jobname, "event": event, "signature": signature})
		}
	}
}

func HandlePing(c *gin.Context) {
	c.String(200, "pong")
}
