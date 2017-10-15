package main

import (
	"context"
	"fmt"
	"os"

	"github.com/lestrrat/go-slack"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	token := os.Getenv("SLACK_API_TOKEN")
	cl := slack.New(token)

	authers, err := cl.Auth().Test().Do(ctx)
	if err != nil {
		fmt.Printf("failed to test authentication: %s\n", err)
		return
	}
	botID := authers.UserID

	listener := NewSlackListener(cl, botID)

	listener.ListenAndResponse()
}
