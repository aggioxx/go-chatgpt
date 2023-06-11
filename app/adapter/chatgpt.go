package adapter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type ChatGptInterface interface {
	SendChat(input string) (string, error)
}

type ResponseStruct struct {
	Created float64  `json:"created"`
	ID      string   `json:"id"`
	Model   string   `json:"model"`
	Object  string   `json:"object"`
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Text         string `json:"text"`
	FinishReason string `json:"finish_reason"`
}

func SendChat(input string) (string, error) {
	apiKey := "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"

	reqBody, err := json.Marshal(map[string]interface{}{
		"temperature": 0,
		"prompt":      input,
		"model":       "text-davinci-003",
		"max_tokens":  400,
	})
	if err != nil {
		fmt.Println("error mounting request body", err.Error())
		return "", err
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println("error to make request", err.Error())
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error to execute request", err.Error())
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("error to close response")
		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error to get response")
		return "", err
	}

	var response ResponseStruct
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("unmarshal error:", err)
		return "", err
	}

	var generatedText string
	if len(response.Choices) > 0 {
		generatedText = response.Choices[0].Text
	}

	return generatedText, nil
}
