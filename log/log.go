package log

import (
	"context"
	"encoding/json"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/crowemi-io/crowemi-go-utils/cloud"
)

type Logger struct {
	GcpClient *cloud.GcpClient
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
	Path      string    `json:"path" omitempty:"true"`
}

func (logger *Logger) Log(message string, level LogLevel, obj any, path string) (string, error) {
	gcp := logger.GcpClient
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, gcp.Config.ProjectID)
	if err != nil {
		return "", err
	}
	defer client.Close()

	logMessage := LogMessage{
		CreatedAt: time.Now(),
		App:       gcp.App,
		Message:   message,
		Level:     level.String(),
		Obj:       obj,
		Path:      path,
	}
	m, err := json.Marshal(logMessage)
	if err != nil {
		return "", err
	}

	topic := client.Topic(gcp.Config.Topics["log"])
	result := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(m),
	})
	// Get the server-generated message ID
	id, err := result.Get(ctx)
	if err != nil {
		return "", err
	}
	return id, nil
}
