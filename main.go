package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	adp "go-chatgpt/app/adapter"
)

func main() {
	lambda.Start(adp.Handler)
}
