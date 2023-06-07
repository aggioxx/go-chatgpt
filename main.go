package main

import (
	chat "go-chatgpt/app/adapter"
)

type Request struct {
	Message string `json:"message"`
}

var messageInput string

func main() {
	//lambda.Start(lambdaHandler)
	callChat()
}

func callChat() string {
	input := "defina cristiano ronaldo em 1 palavra"
	sendChat, err := chat.SendChat(input)
	if err != nil {
		return "error"
	}
	return sendChat
}

//func lambdaHandler(ctx context.Context, request events.APIGatewayProxyRequest) {
//	message := request.Body
//	messageInput = message
//}
