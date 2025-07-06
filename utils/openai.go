package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Request/Response format for OpenAI API
type OpenAIRequest struct {
	Model    string         `json:"model"`
	Messages []OpenAIMessage `json:"messages"`
}

type OpenAIMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIResponse struct {
	Choices []struct {
		Message OpenAIMessage `json:"message"`
	} `json:"choices"`
	Error *struct {
		Message string `json:"message"`
		Type    string `json:"type"`
	} `json:"error,omitempty"`
}

// CallOpenAIWithName generates a response based on name and question
func CallOpenAIWithName(name, question string) (string, error) {
	prompt := fmt.Sprintf("Hi, my name is %s. I have a finance question: %s. Can you help me understand it in a simple way?", name, question)
	return callOpenAI(prompt)
}

// CallOpenAIRecommendations generates a list of recommendations based on a previous result
func CallOpenAIRecommendations(name, prev string) (string, error) {
	prompt := fmt.Sprintf("Hey %s! Based on this context: '%s', give me 5 specific financial recommendations to improve my situation.", name, prev)
	return callOpenAI(prompt)
}

// callOpenAI sends the actual HTTP request to OpenAI API
func callOpenAI(prompt string) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", errors.New("OpenAI API key not set in environment")
	}

	payload := OpenAIRequest{
		Model: "gpt-3.5-turbo",
		Messages: []OpenAIMessage{
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		data, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("OpenAI API error: %s", string(data))
	}

	var result OpenAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if result.Error != nil {
		return "", errors.New(result.Error.Message)
	}

	if len(result.Choices) == 0 {
		return "", errors.New("no response from OpenAI")
	}

	return result.Choices[0].Message.Content, nil
}
