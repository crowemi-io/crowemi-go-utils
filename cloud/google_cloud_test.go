package cloud

import (
	"testing"

	"github.com/crowemi-io/crowemi-go-utils/config"
)

func setup() *GoogleCloudClient {
	config, _ := config.Bootstrap[config.GoogleCloud]("../.secret/config-gcp.json")
	client := &GoogleCloudClient{
		Config: config,
	}
	return client
}

func TestLogMessage(t *testing.T) {
	c := setup()
	id, err := c.Log(
		"Hello 👋 from crowemi-go-utils!",
		INFO,
		nil,
		"google_cloud_test.TestLogMessage",
	)
	if err != nil {
		t.Errorf("Failed to send message: %e", err)
	}
	if id == "" {
		t.Errorf("Message ID: %s", id)
	}
}
