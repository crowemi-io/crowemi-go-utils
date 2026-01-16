package firestore

import (
	"context"
	"testing"

	"github.com/crowemi-io/crowemi-go-utils/config"
	"github.com/crowemi-io/crowemi-go-utils/db"
)

type obj struct {
	Name   string `firestore:"name,omitempty"`
	Salary int64  `firestore:"salary,omitempty"`
}

func setup() (*Client, error) {
	config, err := config.Bootstrap[config.GoogleCloud]("../../.secret/config-google-cloud.json")
	c := Client{
		Config: config,
	}
	return &c, err
}

func TestFirestoreConnect(t *testing.T) {
	c, err := setup()
	if err != nil {
		t.Errorf("Failed to setup Firestore: %v", err)
	}
	firestoreClient, err := c.Connect(context.TODO())
	if firestoreClient != nil {
		t.Logf("Connected to Firestore successfully")
	} else {
		if err != nil {
			t.Logf("Failed to connect to Firestore: %v", err)
		}
		t.Errorf("Failed to connect to Firestore: %v", c)
	}
	defer firestoreClient.Close()
}

func TestFirestoreGetOne(t *testing.T) {
	c, err := setup()
	if err != nil {
		t.Errorf("Failed to setup Firestore: %v", err)
	}
	firestoreClient, err := c.Connect(context.TODO())
	if err != nil {
		t.Errorf("Failed to connect to Firestore: %v", err)
	}
	defer firestoreClient.Close()

	doc, err := GetOne[obj](context.TODO(), firestoreClient, "test", "ZF52AXmmOMRfTamE4dEv")
	if err != nil {
		t.Errorf("Failed to get one document from Firestore: %v", err)
	}
	if doc != nil {
		t.Logf("Document name: %s", doc.Name)
	}
	// t.Logf("Document ID: %s", doc.Ref.ID)
	// t.Logf("Document data: %v", doc.Data())
}
func TestFirestoreGetMany(t *testing.T) {
	c, err := setup()
	if err != nil {
		t.Errorf("Failed to setup Firestore: %v", err)
	}
	firestoreClient, err := c.Connect(context.TODO())
	if err != nil {
		t.Errorf("Failed to connect to Firestore: %v", err)
	}
	defer firestoreClient.Close()

	f := []db.Filter{}
	f = append(f, db.Filter{Field: "date", Operator: "==", Value: "January 12, 2026"})

	doc, err := GetMany[obj](context.TODO(), firestoreClient, "test", f)
	if err != nil {
		t.Errorf("Failed to get one document from Firestore: %v", err)
	}
	for _, d := range *doc {
		t.Logf("Document name: %s", d.Name)
	}
}
func TestFireststoreInsertOne(t *testing.T) {
	c, err := setup()
	if err != nil {
		t.Errorf("Failed to setup Firestore: %v", err)
	}
	firestoreClient, err := c.Connect(context.TODO())
	if err != nil {
		t.Errorf("Failed to connect to Firestore: %v", err)
	}
	defer firestoreClient.Close()

	o := obj{
		Name:   "John Doe",
		Salary: 1000000,
	}

	ret, res, err := InsertOne(context.TODO(), firestoreClient, "test", o)
	if err != nil {
		t.Logf("Failed to insert one document to Firestore: %v", err)
	}
	t.Logf("Document name: %s", ret.ID)
	t.Logf("Document name: %s", res)
}
func TestFirestoreUpdateOne(t *testing.T) {
	c, err := setup()
	if err != nil {
		t.Errorf("Failed to setup Firestore: %v", err)
	}
	firestoreClient, err := c.Connect(context.TODO())
	if err != nil {
		t.Errorf("Failed to connect to Firestore: %v", err)
	}
	defer firestoreClient.Close()

	updates := []Update{}
	f := Update{
		Path:  "name",
		Value: "John Doe",
	}
	updates = append(updates, f)
	res, err := UpdateOne(context.TODO(), firestoreClient, "test", "7F31jh6xszH5331ICwUN", updates)
	if err != nil {
		t.Logf("failed updating one: %s", err)
	}
	t.Logf("Document name: %s", res)
}
func TestFirestoreDeleteOne(t *testing.T) {
	c, err := setup()
	if err != nil {
		t.Errorf("Failed to setup Firestore: %v", err)
	}
	firestoreClient, err := c.Connect(context.TODO())
	if err != nil {
		t.Errorf("Failed to connect to Firestore: %v", err)
	}
	defer firestoreClient.Close()

	res, err := DeleteOne(context.TODO(), firestoreClient, "test", "7F31jh6xszH5331ICwUN")
	if err != nil {
		t.Logf("failed updating one: %s", err)
	}
	t.Logf("Document name: %s", res)
}
