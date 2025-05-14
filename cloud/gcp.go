package cloud

import (
	"context"

	"cloud.google.com/go/storage"
	"github.com/crowemi-io/crowemi-go-utils/config"
)

type GcpClient struct {
	App    string
	Config *config.GoogleCloud
	// Add GCP components
	CloudStorage *CloudStorage
}

func NewGcpClient(app string, config *config.GoogleCloud) (*GcpClient, error) {
	storageClient, err := storage.NewClient(context.TODO())
	if err != nil {
		return nil, err
	}

	bucket := storageClient.Bucket(app)

	cloudStorage := CloudStorage{
		Bucket:        bucket,
		StorageClient: storageClient,
	}
	client := GcpClient{
		App:          app,
		Config:       config,
		CloudStorage: &cloudStorage,
	}
	return &client, nil
}
