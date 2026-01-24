package ai

import (
	"context"

	"github.com/ollama/ollama/api"
)

var client *api.Client

func Init() error {
	c, err := api.ClientFromEnvironment()
	if err != nil {
		return err
	}

	client = c

	return nil
}

func True() *bool {
	b := true
	return &b
}

func False() *bool {
	b := false
	return &b
}

func Generate(ctx context.Context, p string) string {
	req := &api.GenerateRequest{
		Model:     "gpt-oss",
		Prompt:    p,
		KeepAlive: &api.Duration{Duration: 120},

		Stream: False(),
	}

	var res string

	if err := client.Generate(ctx, req, func(resp api.GenerateResponse) error {
		res = resp.Response
		return nil
	}); err != nil {
		return err.Error()
	}

	return res
}
