package main

import (
	"context"

	"github.com/joho/godotenv"
	"patrick.com/render-atl-hackathon/chatgpt"
)

func main() {

	godotenv.Load(".env")
	chatgpt.SummarizeReview(context.TODO())
}
