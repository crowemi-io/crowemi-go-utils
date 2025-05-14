package log

import (
	"testing"

	"github.com/crowemi-io/crowemi-go-utils/cloud"
	"github.com/crowemi-io/crowemi-go-utils/config"
)

func setup() *cloud.GcpClient {
	config, _ := config.Bootstrap[config.GoogleCloud]("../.secret/config-gcp.json")
	client := &cloud.GcpClient{
		Config: config,
	}
	return client
}

func TestLogMessage(t *testing.T) {
	c := setup()
	logger := Logger{GcpClient: c}
	id, err := logger.Log(
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
