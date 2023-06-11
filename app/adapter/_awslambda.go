package adapter

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

type MyInputType struct {
	Text string `json:"text"`
}

type MyOutputType struct {
	Text string `json:"text"`
}

func handleRequest(ctx context.Context, input events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Parse the request body into MyInputType struct
	var myInput MyInputType
	err := json.Unmarshal([]byte(input.Body), &myInput)
	if err != nil {
		fmt.Println("error parsing request body:", err)
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}

	// Invoke the SendChat function from the adapter package
	generatedText, err := SendChat(myInput.Text)
	if err != nil {
		fmt.Println("error calling SendChat:", err)
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}

	// Create the response object
	response := MyOutputType{
		Text: generatedText,
	}

	// Convert the response to JSON
	responseBody, err := json.Marshal(response)
	if err != nil {
		fmt.Println("error marshaling response:", err)
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}

	// Create the API Gateway response
	apiResponse := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(responseBody),
	}

	return apiResponse, nil
}
