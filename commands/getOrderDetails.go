package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Structs to match the request and response JSON structures
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type RequestPayload struct {
	Messages []Message `json:"messages"`
	Model    string    `json:"model"`
}

type ResponseMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Choice struct {
	Index   int             `json:"index"`
	Message ResponseMessage `json:"message"`
}

type ResponsePayload struct {
	ID      string   `json:"id"`
	Choices []Choice `json:"choices"`
}

// Function to call the API and capture the response
func callAPI(messageContent string) (string, error) {
	// Create the payload with the dynamic content
	payload := RequestPayload{
		Messages: []Message{
			{
				Role:    "user",
				Content: messageContent, // Use the passed argument as content
			},
		},
		Model: "llama3-8b-8192",
	}

	// Convert the payload to JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal payload: %v", err)
	}

	// Send the POST request
	url := "https://api.groq.com/openai/v1/chat/completions" // API URL
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	// Set headers
	apiKey := os.Getenv("GROQ_API_KEY") // Assuming the API key is stored in an environment variable
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("Content-Type", "application/json")

	// Create an HTTP client and perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	// Unmarshal the response JSON into ResponsePayload struct
	var responsePayload ResponsePayload
	err = json.Unmarshal(body, &responsePayload)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %v", err)
	}

	// Extract the content from the response
	if len(responsePayload.Choices) > 0 {
		return responsePayload.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no choices found in response")
}

// func main() {
// 	response, err := callAPI()
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	fmt.Println("Response Content:", response)
// }
