package cloud

import (
	"context"
	"encoding/json"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/crowemi-io/crowemi-go-utils/config"
)

type GoogleCloudClient struct {
	Config *config.GoogleCloud
}

type LogLevel int

const (
	INFO LogLevel = iota
	ERROR
	WARNING
	DEBUG
)

func (d LogLevel) String() string {
	level := []string{"INFO", "ERROR", "WARNING", "DEBUG"}
	return level[d]
}

type LogMessage struct {
	CreatedAt time.Time `json:"created_at" omitempty:"true"`
	App       string    `json:"app" omitempty:"true"`
	Message   string    `json:"message" omitempty:"true"`
	Level     string    `json:"level" omitempty:"true"`
	Obj       any       `json:"obj" omitempty:"true"`
	Session   string    `json:"session" omitempty:"true"`
	Path      string    `json:"path" omitempty:"true"`
}

func (logMessage *LogMessage) Log(projectID string, topicID string) (string, error) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return "", err
	}
	defer client.Close()

	message, err := json.Marshal(logMessage)
	if err != nil {
		return "", err
	}
	topic := client.Topic(topicID)
	result := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(message),
	})
	// Get the server-generated message ID
	id, err := result.Get(ctx)
	if err != nil {
		return "", err
	}
	return id, nil
}
