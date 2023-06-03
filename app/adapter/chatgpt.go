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

func SendChat(input string) (string, error) {
	apiKey := "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"

	reqBody, err := json.Marshal(map[string]interface{}{
		"temperature": 0,
		"prompt":      input,
		"model":       "text-davinci-003",
		"max_tokens":  40,
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

	var response interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("unmarshal error:", err)
		return "", err
	}
	fmt.Println(body, response)
	return string(body), nil
}
