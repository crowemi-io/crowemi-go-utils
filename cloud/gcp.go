package cloud

import "github.com/crowemi-io/crowemi-go-utils/config"

type GcpClient struct {
	App     string
	Session string
	Config  *config.GoogleCloud
	// Add GCP components
	CloudStorage *CloudStorage
}
