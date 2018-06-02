package main

import (
	"context"
	"log"
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
		log.Fatalf("failed to test authentication: %s\n", err)
	}
	botID := authers.UserID

	listener := NewSlackListener(cl, botID)

	listener.ListenAndResponse()
}
