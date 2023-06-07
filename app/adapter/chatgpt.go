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
	Created float64 `json:"created"`
	ID      string  `json:"id"`
	Model   string  `json:"model"`
	Object  string  `json:"object"`
	Text    string  `json:"text"`
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

	var responseMap map[string]interface{}
	err = json.Unmarshal(body, &responseMap)
	if err != nil {
		fmt.Println("unmarshal error:", err)
		return "", err
	}

	jsonBytes, err := json.MarshalIndent(responseMap, "", "  ")
	if err != nil {
		fmt.Println("json marshal error:", err)
		return "", err
	}

	fmt.Println(responseMap["text"])
	fmt.Println("")
	fmt.Println(string(jsonBytes))
	return string(body), nil

}
