package genai

import (
	"context"

	"github.com/crowemi-io/crowemi-go-utils/config"
	"google.golang.org/genai"
)

type Client struct {
	Config *config.GoogleCloud
}

func (c *Client) Connect(ctx context.Context) (*genai.Client, error) {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: c.Config.GenAI.ApiKey,
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (c *Client) Generate(ctx context.Context, content string) (string, error) {
	client, err := c.Connect(ctx)
	if err != nil {
		return "", err
	}
	parts := []*genai.Part{{Text: content}}
	response, err := client.Models.GenerateContent(ctx, c.Config.GenAI.Model, []*genai.Content{{Parts: parts}}, nil)
	if err != nil {
		return "", err
	}
	return response.Text(), nil
}
