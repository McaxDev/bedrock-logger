package main

import (
	"time"
)

type ChatMessage struct {
	Username string
	Source   string
	Content  string
	Time     time.Time
}

var ChatMessages struct {
	Messages []ChatMessage
	Iterator int
}

func InitChatMessages() {
	ChatMessages.Messages = make([]ChatMessage, config.ChatLength)
}
