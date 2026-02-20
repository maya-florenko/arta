package ollama

import (
	"context"
	_ "embed"

	"github.com/ollama/ollama/api"
)

var ollama *api.Client

//go:embed prompt.md
var s string

func New() error {
	c, err := api.ClientFromEnvironment()
	if err != nil {
		return err
	}

	ollama = c

	return nil
}

func Generate(ctx context.Context, p string) string {
	var response string

	req := &api.GenerateRequest{
		Model:  "gpt-oss:120b-cloud",
		Prompt: p,
		System: s,
		Stream: new(bool),
		Think:  &api.ThinkValue{Value: "medium"},
		Options: map[string]any{
			"temperature": 0.6,
		},
	}

	res := func(resp api.GenerateResponse) error {
		response = resp.Response
		return nil
	}

	err := ollama.Generate(ctx, req, res)
	if err != nil {
		return err.Error()
	}

	return response
}
