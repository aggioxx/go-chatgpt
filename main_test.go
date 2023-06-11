package main

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandleRequest(t *testing.T) {

	event := events.APIGatewayProxyRequest{
		Body: `{"data": "Qual a profissão do lionel messi?"}`,
	}

	response, err := handleRequest(event)
	if err != nil {
		t.Fatalf("handleRequest returned an error: %s", err.Error())
	}

	fmt.Println("Response:", response)

	var responseBody MyOutputType
	err = json.Unmarshal([]byte(response.Body), &responseBody)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %s", err.Error())
	}

	fmt.Println("Generated Text:", responseBody.Text)
	assert.NotNil(t, responseBody.Text)
}
