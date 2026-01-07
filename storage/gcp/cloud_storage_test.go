package gcp

import (
	"context"
	"encoding/json"
	"testing"

	"cloud.google.com/go/storage"
)

func setUp() (*Client, error) {
	storageClient, err := storage.NewClient(context.TODO())
	if err != nil {
		return nil, err
	}
	bucket := storageClient.Bucket("crowemi-trades")
	cloudStorage := Client{
		Bucket:        bucket,
		StorageClient: storageClient,
	}
	return &cloudStorage, nil
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

type Hello struct {
	Hello string
	World string
}

func TestRead(t *testing.T) {
	c, err := setUp()
	file := "test/test.json"
	if err != nil {
		t.Fatal("failed to create storage client")
	}

	data, err := c.Read(file)
	if err != nil {
		t.Fatalf("failed to reade file %s", file)
	}

	h := &Hello{}
	err = json.Unmarshal(data, h)
	if err != nil {
		t.Fatalf("failed to unmarshal json")
	}

	if h.Hello != "Hello" {
		t.Fail()
	}
	if h.World != "World" {
		t.Fail()
	}
}
