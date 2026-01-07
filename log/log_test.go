package log

import (
	"testing"

	"github.com/crowemi-io/crowemi-go-utils/config"
	"github.com/crowemi-io/crowemi-go-utils/storage/cloud_storage"
)

func setup() *cloud_storage.Client {
	config, _ := config.Bootstrap[config.GoogleCloud]("../.secret/config-gcp.json")
	client := &cloud_storage.Client{
		Config: config,
	}
	return client
}

func TestLogMessage(t *testing.T) {
	c := setup()
	logger := Logger{CloudStorage: c}
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
