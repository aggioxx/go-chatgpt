package main

import (
	chat "go-chatgpt/app/adapter"
)

/*
	type chatImpl struct {
		chat *adapter.ChatGptInterface
	}

	func NewChatImpl(chat *adapter.ChatGptInterface) *chatImpl {
		return &chatImpl{
			chat: chat,
		}
	}
*/
func callChat() string {
	input := "qual a profissão do lionel messi?"
	sendChat, err := chat.SendChat(input)
	if err != nil {
		return ""
	}
	return sendChat
}

func main() {
	callChat()
}
