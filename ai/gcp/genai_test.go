package genai

import (
	"context"
	"testing"

	"github.com/crowemi-io/crowemi-go-utils/config"
)

func setup() (*Client, error) {
	config, err := config.Bootstrap[config.GoogleCloud]("../../.secret/config-google-cloud.json")
	c := Client{
		Config: config,
	}
	return &c, err
}

func TestGenerate(t *testing.T) {
	c, err := setup()
	if err != nil {
		t.Fatal(err)
	}
	content, err := c.Generate(context.Background(), "Hello robot, from world!")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(content)
}
