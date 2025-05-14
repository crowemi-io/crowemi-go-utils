package cloud

import (
	"context"
	"encoding/json"
	"testing"

	"cloud.google.com/go/storage"
)

func setUp() (*GcpClient, error) {
	storageClient, err := storage.NewClient(context.TODO())
	if err != nil {
		return nil, err
	}
	bucket := storageClient.Bucket("crowemi-trades")
	cloudStorage := CloudStorage{
		Bucket:        bucket,
		StorageClient: storageClient,
	}
	client := GcpClient{
		CloudStorage: &cloudStorage,
	}
	return &client, nil
}

func TestWrite(t *testing.T) {
	c, err := setUp()
	if err != nil {
		t.Fatal("Failed to create storage client")
	}

	d := struct {
		Hello string
		World string
	}{
		Hello: "Hello",
		World: "World",
	}
	b, _ := json.Marshal(d)
	c.Write("test/test.json", b)

}
