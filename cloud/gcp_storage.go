package cloud

import (
	"context"

	"cloud.google.com/go/storage"
)

type GcpStorageClient struct {
	Bucket        *storage.BucketHandle
	StorageClient *storage.Client
}

func (c *GcpStorageClient) Write(prefix string, contents []byte) error {
	object := c.Bucket.Object(prefix)
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
func (client *GcpStorageClient) Read() {}
