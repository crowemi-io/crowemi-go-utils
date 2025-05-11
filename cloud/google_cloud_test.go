package cloud

import (
	"testing"
	"time"

	"github.com/crowemi-io/crowemi-go-utils/config"
)

func setup() *config.GoogleCloud {
	config, _ := config.Bootstrap[config.GoogleCloud]("../.secret/config-gcp.json")
	return config
}

func TestLogMessage(t *testing.T) {
	c := setup()
	m := LogMessage{
		CreatedAt: time.Now(),
		App:       "crowemi-go-utils",
		Message:   "Hello 👋, from crowemi-go-utils",
		Level:     INFO.String(),
		Path:      "google_cloud_test.TestLogMessage",
	}
	id, err := m.Log(c.ProjectId, c.Topic)
	if err != nil {
		t.Errorf("Failed to send message: %e", err)
	}
	if id == "" {
		t.Errorf("Message ID: %s", id)
	}
}
