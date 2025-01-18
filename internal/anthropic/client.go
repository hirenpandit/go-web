package anthropic

import (
	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file %v\n", err)
	}
}

func NewClient() *anthropic.Client {
	return anthropic.NewClient(option.WithAPIKey(os.Getenv("ANTHROPIC_API_KEY")))
}
