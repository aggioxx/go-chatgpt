package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	apiKey := "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"
	input := "quem é cristiano ronaldo?"

	reqBody, err := json.Marshal(map[string]interface{}{
		"temperature": 0,
		"prompt":      input,
		"model":       "text-davinci-003",
		"max_tokens":  4000,
	})
	if err != nil {
		fmt.Println("Erro ao montar o body", err.Error())
		return
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println("Erro na request", err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao executar a request", err.Error())
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Erro ao ler a response")
		return
	}
	fmt.Println(string(body))
}
