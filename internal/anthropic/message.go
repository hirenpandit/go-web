package anthropic

import (
	"context"
	"github.com/anthropics/anthropic-sdk-go"
	"log"
)

func SendMessage(query string) string {
	client := NewClient()
	message, err := client.Messages.New(context.TODO(), anthropic.MessageNewParams{
		Model:     anthropic.F(anthropic.ModelClaude3_5SonnetLatest),
		MaxTokens: anthropic.F(int64(1024)),
		Messages: anthropic.F([]anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock(query)),
		}),
	})

	if err != nil {
		log.Printf("error sending message: %v", err)
	}
	return message.Content[0].Text
}

func SendMessageWithStreamResponse(query string) {
	log.Println("sending query to anthropic >>", query)
	client := NewClient()
	stream := client.Messages.NewStreaming(context.TODO(), anthropic.MessageNewParams{
		Model:     anthropic.F(anthropic.ModelClaude3_5SonnetLatest),
		MaxTokens: anthropic.F(int64(1024)),
		Messages: anthropic.F([]anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock(query)),
		}),
	})

	message := anthropic.Message{}

	for stream.Next() {
		event := stream.Current()
		err := message.Accumulate(event)
		if err != nil {
			log.Printf("error sending message: %v", err)
		}
		switch delta := event.Delta.(type) {
		case anthropic.ContentBlockDeltaEventDelta:
			if delta.Text != "" {
				log.Printf(">> %s", delta.Text)
			}
		}
	}

}
