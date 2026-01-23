package ai

import (
	"context"
	"log"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

func Init(c context.Context) string {
	client := openai.NewClient(
		option.WithBaseURL("http://127.0.0.1:8080/v1"),
	)

	completion, err := client.Chat.Completions.New(c, openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage("meow"),
		},
		Model: openai.ChatModelGPT4o,
	})
	if err != nil {
		log.Fatal(err)
	}

	return completion.Choices[0].Message.Content
}
