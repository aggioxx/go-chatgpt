package adapter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

// RateLimiter defines a rate limiter based on the token bucket algorithm.
type RateLimiter struct {
	mu           sync.Mutex
	tokens       int
	maxTokens    int
	refillAmount int
	refillPeriod time.Duration
	lastRefill   time.Time
}

// NewRateLimiter creates a new rate limiter with the specified parameters.
func NewRateLimiter(maxTokens, refillAmount int, refillPeriod time.Duration) *RateLimiter {
	return &RateLimiter{
		tokens:       maxTokens,
		maxTokens:    maxTokens,
		refillAmount: refillAmount,
		refillPeriod: refillPeriod,
		lastRefill:   time.Now(),
	}
}

// Acquire acquires tokens from the rate limiter, blocking until tokens are available.
func (rl *RateLimiter) Acquire(n int) {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	rl.refill()

	for rl.tokens < n {
		time.Sleep(10 * time.Millisecond)
		rl.refill()
	}

	rl.tokens -= n
}

// refill refills tokens in the rate limiter if the refill period has passed.
func (rl *RateLimiter) refill() {
	now := time.Now()
	elapsed := now.Sub(rl.lastRefill)
	refills := int(elapsed / rl.refillPeriod)

	if refills > 0 {
		rl.tokens = rl.tokens + rl.refillAmount*refills
		if rl.tokens > rl.maxTokens {
			rl.tokens = rl.maxTokens
		}
		rl.lastRefill = now
	}
}

func SendChat(input string) (string, error) {
	apiKey := "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"
	rateLimiter := NewRateLimiter(5, 5, 1*time.Second) // Adjust rate limits as per OpenAI guidelines

	for {
		rateLimiter.Acquire(1) // Acquire 1 token before making the API request

		reqBody, err := json.Marshal(map[string]interface{}{
			"temperature": 0,
			"prompt":      input,
			"model":       "text-davinci-003",
			"max_tokens":  40,
		})
		if err != nil {
			return "", err
		}

		req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", bytes.NewBuffer(reqBody))
		if err != nil {
			return "", err
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+apiKey)

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			return "", err
		}

		if res.StatusCode == http.StatusOK {
			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				return "", err
			}

			var response struct {
				Choices []struct {
					Text string `json:"text"`
				} `json:"choices"`
			}

			err = json.Unmarshal(body, &response)
			if err != nil {
				return "", err
			}

			if len(response.Choices) > 0 {
				return response.Choices[0].Text, nil
			}
		} else if res.StatusCode == http.StatusTooManyRequests {
			retryAfter := res.Header.Get("Retry-After")
			retryAfterDuration, _ := time.ParseDuration(retryAfter)
			time.Sleep(retryAfterDuration)
		} else {
			return "", fmt.Errorf("API request failed with status: %s", res.Status)
		}
	}
}
