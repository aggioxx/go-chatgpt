package main

import (
	chat "go-chatgpt/app/adapter"
)

func callChat() string {
	input := "what is a nil pointer exception in golang"
	sendChat, err := chat.SendChat(input)
	if err != nil {
		return ""
	}
	return sendChat
}

func main() {
	callChat()
}
