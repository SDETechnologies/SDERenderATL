package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

func GetClient() *openai.Client {
	godotenv.Load()
	token := os.Getenv("OPEN_API_TOKEN")

	return openai.NewClient(token)
}

type ChatGPTResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		Logprobs     any    `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	SystemFingerprint string `json:"system_fingerprint"`
}

func SummarizeTone(ctx context.Context) error {
	text := `Pro : Direct access from Airport to ATL downtown area.

Con : Signage in downtown area are lacking in the platform to which street exits.

Cheap and reliable way from airport to downtown. Entrance from airport is right next to the baggage claim, buy ticket and go up level to ride the MARTA into town. Bypass the ticket line by purchasing the ticket online... it is the busiest airport on the planet.

Got off at Peachtree and it was a maze without much signage, I end up walking out of the North exit and roam on the city streets southbound to the south exit. Station was quite deserted at night, kind of creepy.`
	apiKey := os.Getenv("OPEN_API_TOKEN")
	client := resty.New()
	apiEndpoint := "https://api.openai.com/v1/chat/completions"

	message1 := map[string]string{"role": "system",
		"content": "You are a data analsyst assistant that provides concise details about reviews of a transportation system for a large city. Avoid details and isntead generalize. Here is an example response, do not include any formatting, only json {\"Good Things\": [\"The trains were on time\", \"They were clean\"], \"Bad Things\": [\"I wasn't sure what side of the train to get off of\", \"The exits getting out of the station were hard to follow\"]}",
	}
	message2 := map[string]string{"role": "user", "content": fmt.Sprintf("here is my review of my marta trip:\n\n%s", text)}
	message := []map[string]string{message1, message2}

	response, err := client.R().
		SetAuthToken(apiKey).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]any{
			"model":      "gpt-4o",
			"messages":   message,
			"max_tokens": 500,
		}).
		Post(apiEndpoint)

	if err != nil {
		panic(err)
	}

	res := ChatGPTResponse{}

	err = json.Unmarshal(response.Body(), &res)
	if err != nil {
		panic(err)
	}

	fmt.Println("++++++++=")
	fmt.Println(res.Choices[0].Message.Content)
	fmt.Println("++++++++=")
	final, err := json.MarshalIndent(res.Choices[0].Message.Content, "", "  ")

	if err != nil {
		panic(err)
	}
	fmt.Println(string(final))

	return nil
}
