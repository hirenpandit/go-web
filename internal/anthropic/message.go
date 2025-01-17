package anthropic

import (
	"context"
	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	log.Printf("loading environment variables")
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("error loading .env file %v", err)
	}
}

func SendMessage() string {
	client := anthropic.NewClient(
		option.WithAPIKey(os.Getenv("ANTHROPIC_API_KEY")),
	)

	message, err := client.Messages.New(context.TODO(), anthropic.MessageNewParams{
		Model:     anthropic.F(anthropic.ModelClaude3_5SonnetLatest),
		MaxTokens: anthropic.F(int64(1024)),
		Messages: anthropic.F([]anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock("What is a quaternion?")),
		}),
	})

	if err != nil {
		log.Printf("error sending message: %v", err)
	}
	return message.Content[0].Text
}
