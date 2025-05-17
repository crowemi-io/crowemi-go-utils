package cloud

import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
)

type CloudStorage struct {
	Bucket        *storage.BucketHandle
	StorageClient *storage.Client
}

func (c *GcpClient) Write(prefix string, contents []byte) (int, error) {
	object := c.CloudStorage.Bucket.Object(prefix)
	writer := object.NewWriter(context.Background())

	ret, err := writer.Write(contents)
	if err != nil {
		return -1, err
	}
	defer func() {
		closeErr := writer.Close()
		if closeErr != nil {
			if err != nil {
				err = fmt.Errorf("%w; additionally, failed to close reader: %v", err, closeErr)
			} else {
				err = fmt.Errorf("failed to close reader: %w", closeErr)
			}
		}
	}()

	return ret, err
}
func (c *GcpClient) Read(prefix string) ([]byte, error) {
	object := c.CloudStorage.Bucket.Object(prefix)
	reader, err := object.NewReader(context.Background())
	if err != nil {
		return nil, err
	}
	defer func() {
		closeErr := reader.Close()
		if closeErr != nil {
			if err != nil {
				err = fmt.Errorf("%w; additionally, failed to close reader: %v", err, closeErr)
			} else {
				err = fmt.Errorf("failed to close reader: %w", closeErr)
			}
		}
	}()

	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return data, err
}
