package chatgpt

import (
	"context"
	"encoding/json"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
	"patrick.com/render-atl-hackathon/db"
)

func GetClient() *openai.Client {
	godotenv.Load()
	token := os.Getenv("OPEN_API_TOKEN")

	return openai.NewClient(token)
}

type RequestBody struct {
	ResponseFormat map[string]string `json:"response_format"`
	Model          string            `json:"model"`
	Messages       any               `json:"messages"`
	MaxToken       int               `json:"max_token"`
}

func SummarizeReview(ctx context.Context, review string) (db.Feedback, error) {
	godotenv.Load(".env")
	apiKey := os.Getenv("OPEN_API_TOKEN")
	client := resty.New()

	response, err := client.R().
		SetAuthToken(apiKey).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]any{
			"response_format": map[string]string{"type": "json_object"},
			"model":           "gpt-4o",
			"messages":        []map[string]string{map[string]string{"role": "system", "content": inititalPrompt}, map[string]string{"role": "user", "content": directions}, map[string]string{"role": "user", "content": review}},
			"max_tokens":      1000,
		}).
		Post(apiEndpoint)

	if err != nil {
		panic(err)
	}

	chatGPTResponse := ChatGPTResponse{}

	err = json.Unmarshal(response.Body(), &chatGPTResponse)

	if err != nil {
		panic(err)
	}

	feedback := db.Feedback{}

	err = json.Unmarshal([]byte(chatGPTResponse.Choices[0].Message.Content), &feedback)

	if err != nil {
		panic(err)
	}

	return feedback, nil
}
