package cloud

import "github.com/crowemi-io/crowemi-go-utils/config"

type GcpClient struct {
	App           string
	Session       string
	Config        *config.GoogleCloud
	StorageClient *GcpStorageClient
}
