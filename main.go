package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/webhooks/v6/github"
)

const (
	path = "/webhooks"
)

func main() {
	hook, _ := github.New(github.Options.Secret("6217512Fang"))
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.POST(path, func(c *gin.Context) {
		payload, err := hook.Parse(c.Request, github.ReleaseEvent, github.PullRequestEvent)
		if err != nil {
			fmt.Println("error is: " + err.Error())
		}
		switch payload.(type) {

		case github.PushPayload:
			fmt.Println(11111)
		case github.ReleasePayload:
			release := payload.(github.ReleasePayload)
			// Do whatever you want from here...
			fmt.Printf("%+v", release)

		case github.PullRequestPayload:
			pullRequest := payload.(github.PullRequestPayload)
			// Do whatever you want from here...
			fmt.Printf("%+v", pullRequest)
		}
	})
	router.Run(":3000")
}
