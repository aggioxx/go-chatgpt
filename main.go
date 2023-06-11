package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"go-chatgpt/app/adapter"
	"strings"
)

type MyInputType struct {
	Data string `json:"data"`
}

func handleRequest(input MyInputType) (string, error) {

	generatedText, err := adapter.SendChat(input.Data)
	if err != nil {
		fmt.Println("error calling SendChat:", err)
		return "", err
	}

	generatedText = strings.TrimPrefix(generatedText, "\n\n")
	return generatedText, nil
}

func main() {
	lambda.Start(handleRequest)
}
