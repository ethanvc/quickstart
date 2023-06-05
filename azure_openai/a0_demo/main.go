package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func toJsonStr(v any) string {
	buf, _ := json.Marshal(v)
	return string(buf)
}

type CompletionsReq struct {
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func main() {
	apiKey := os.Getenv("AZURE_OPENAI_KEY")
	endpoint := os.Getenv("AZURE_OPENAI_ENDPOINT")
	url := fmt.Sprintf("%s/openai/deployments/gpt35/chat/completions?api-version=2023-05-15", endpoint)
	var req CompletionsReq
	req.Messages = append(req.Messages,
		Message{
			Role:    "system",
			Content: "you are a helpful assistant.",
		},
		Message{
			Role:    "user",
			Content: "Does Azure OpenAI support customer managed keys?",
		})
	httpReq, err := http.NewRequest(http.MethodPost, url, strings.NewReader(toJsonStr(req)))
	if err != nil {
		panic(err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("api-key", apiKey)
	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		panic(err)
	}
	httpBody, err := io.ReadAll(httpResp.Body)
	fmt.Println("response", string(httpBody))
}
