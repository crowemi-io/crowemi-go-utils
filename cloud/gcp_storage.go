package cloud

import (
	"context"

	"cloud.google.com/go/storage"
)

type CloudStorage struct {
	Bucket        *storage.BucketHandle
	StorageClient *storage.Client
}

func (c *GcpClient) Write(prefix string, contents []byte) error {
	object := c.CloudStorage.Bucket.Object(prefix)
	writer := object.NewWriter(context.Background())

	_, err := writer.Write(contents)
	if err != nil {
		return err
	}
	err = writer.Close()
	if err != nil {
		return err
	}

	return nil
}
func (client *GcpClient) Read() {}
