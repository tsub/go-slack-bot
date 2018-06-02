package main

import (
	"context"
	"log"
	"strings"

	"github.com/lestrrat/go-slack"
	"github.com/lestrrat/go-slack/rtm"
)

type SlackListener struct {
	Client *slack.Client
	BotID string
}

func NewSlackListener(client *slack.Client, botID string) *SlackListener {
	return &SlackListener{
		Client: client,
		BotID: botID,
	}
}

func (s SlackListener) ListenAndResponse() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rtmcl := rtm.New(s.Client)
	go rtmcl.Run(ctx)

	for e := range rtmcl.Events() {
		switch typ := e.Type(); typ {
		case rtm.MessageType:
			if err := s.handleMessageEvent(ctx, e.Data().(*rtm.MessageEvent)); err != nil {
				log.Printf("[ERROR] Failed to handle message: %s\n", err)
			}
		default:
			// Noop
		}
	}
}

func (s SlackListener) handleMessageEvent(ctx context.Context, e *rtm.MessageEvent) error {
	// Only mention
	if !strings.Contains(e.Text, s.BotID) {
		return nil
	}

	log.Printf("Mention event: %s\n", e)

	if strings.Contains(e.Text, "hi") {
		res, err := s.Client.Chat().PostMessage(e.Channel).Text("hi").Do(ctx)
		if err != nil {
			return err
		}

		log.Printf("Chat response: %s\n", res)
	}

	return nil
}
