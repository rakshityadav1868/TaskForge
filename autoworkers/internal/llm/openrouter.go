package llm

import (
	"autoworkers/internal/job"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type OpenRouterClient struct{
	APIKey string
	Client *http.Client
}
type Request struct{
	Model string `json:"model"`
	Messages []Message `json:"messages"`
	MaxTokens  int    `json:"max_tokens"`
}
type Response struct{
	Choices []Choice `json:"choices"`

}
type Choice struct{
	Message  Message `json:"message"`

}
type Message struct {
	Role string `json:"role"`
	Content string `json:"content"`
}

func Constructor(apikey string) *OpenRouterClient{
	return &OpenRouterClient{
		APIKey: apikey,
		Client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (o *OpenRouterClient) Generate(j *job.Job) (string,error){
	
	var response Response

	req := Request{
		Model: j.Model,
		Messages: []Message{
			{
				Role: "user",
				Content: j.Prompt,
			},
		},
		MaxTokens: 100,
	}
	
	body,err :=json.Marshal(req)

	if err!=nil{
		return "",err
	}

	httpReq, err := http.NewRequest("POST","https://openrouter.ai/api/v1/chat/completions",bytes.NewReader(body))
	if err != nil {
		fmt.Printf("LLM Error: %v\n", err)
    return "", err
	}

	httpReq.Header.Set("Authorization", "Bearer "+o.APIKey)
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := o.Client.Do(httpReq)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		data, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("status %s: %s", resp.Status, string(data))
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Printf("LLM Error: %v\n", err)
    return "", err
	}

	if len(response.Choices) == 0 {
		fmt.Printf("LLM Error: %v\n", err)
    return "", fmt.Errorf("no response from model")
	}

	return response.Choices[0].Message.Content,nil



}