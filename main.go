package main

import (
	"context"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")
	//client := GetClient()
	SummarizeTone(context.TODO())

}
