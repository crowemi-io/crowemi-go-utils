package gcp

import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
	"github.com/crowemi-io/crowemi-go-utils/config"
)

type Client struct {
	Config        *config.GoogleCloud
	Bucket        *storage.BucketHandle
	StorageClient *storage.Client
}

func (c *Client) Write(prefix string, contents []byte) (int, error) {
	object := c.Bucket.Object(prefix)
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
func (c *Client) Read(prefix string) ([]byte, error) {
	object := c.Bucket.Object(prefix)
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
