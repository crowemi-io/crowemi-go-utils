package db

import (
	"context"
	"testing"
	"time"

	"github.com/crowemi-io/crowemi-go-utils/config"
)

func setup() *MongoClient {
	config, _ := config.Bootstrap[config.Crowemi]()
	c := MongoClient{}
	c.Connect(context.TODO(), config.DatabaseURI, config.ClientName)
	return &c
}

func TestMongoDBConnect(t *testing.T) {
	p := setup()
	if p != nil {
		t.Logf("Connected to MongoDB successfully")
	} else {
		t.Errorf("Failed to connect to MongoDB: %v", p)
	}
}
func TestMongoDBPing(t *testing.T) {
	c := setup()
	err := c.Ping()
	if err != nil {
		t.Errorf("Failed to ping MongoDB: %v", err)
	}
}
func TestMongoDBGetOne(t *testing.T) {
	c := setup()
	f := map[string]string{"symbol": "AAPL"}
	type symbol struct{ Symbol string }
	result, err := GetOne[symbol](context.TODO(), *c, "symbol", f)
	if err != nil {
		t.Errorf("Failed to get one document: %v", err)
	}

	if result.Symbol != "AAPL" {
		t.Errorf("Expected symbol AAPL, got %s", result.Symbol)
	}
}
func TestMongoDBGetMany(t *testing.T) {
	c := setup()
	f := map[string]string{"symbol": "AAPL"}
	type symbol struct{ Symbol string }
	result, err := GetMany[symbol](context.TODO(), *c, "orders", f)
	if err != nil {
		t.Errorf("Failed to get one document: %v", err)
	}
	// Ensure the result is not empty by checking a relevant field
	if len(result) == 0 {
		t.Errorf("Expected to get at least one document, got %d", len(result))
	}
}
func TestMongoDBInsertOne(t *testing.T) {
	c := setup()
	type symbol struct {
		Symbol string `bson:"symbol"`
	}
	s := symbol{Symbol: "OXY"}
	ret, err := InsertOne(context.TODO(), *c, "symbol", s)
	if err != nil {
		t.Errorf("Failed to insert one document: %v", err)
	}
	if ret == nil {
		t.Errorf("Expected to get an ID, got nil")
	}
}

func TestMongoDBInsertMany(t *testing.T) {
	c := setup()
	type symbol struct {
		Symbol string `bson:"symbol"`
	}
	s := []symbol{
		{Symbol: "OXY1"},
		{Symbol: "AAPL1"},
	}
	ret, err := InsertMany(context.TODO(), *c, "symbol", s)
	if err != nil {
		t.Errorf("Failed to insert many documents: %v", err)
	}
	if ret == nil {
		t.Errorf("Expected to get an ID, got nil")
	} else if len(ret.InsertedIDs) == 0 {
		t.Errorf("Expected to insert one document, got %d", ret.InsertedIDs...)
	}
}
func TestMongoDBUpdateOne(t *testing.T) {
	c := setup()
	type symbol struct {
		Symbol    string    `bson:"symbol"`
		CreatedAt time.Time `bson:"created_at"`
	}
	s := symbol{Symbol: "OXY", CreatedAt: time.Now()}
	f := map[string]string{"symbol": "OXY"}
	ret, err := UpdateOne(context.TODO(), *c, "symbol", f, s)
	if err != nil {
		t.Errorf("Failed to update one document: %v", err)
	}
	if ret == nil {
		t.Errorf("Expected to get an ID, got nil")
	}
}
func TestMongoDBUpdateMany(t *testing.T) {}
func TestMongoDBDeleteOne(t *testing.T) {
	c := setup()
	f := map[string]string{"symbol": "OXY"}
	ret, err := DeleteOne(context.TODO(), *c, "symbol", f)
	if err != nil {
		t.Errorf("Failed to delete one document: %v", err)
	}
	if ret == nil {
		t.Errorf("Expected to get an ID, got nil")
	} else if ret.DeletedCount == 0 {
		t.Errorf("Expected to delete one document, got %d", ret.DeletedCount)
	}
}
func TestMongoDBDeleteMany(t *testing.T) {
	c := setup()
	f := map[string]string{"symbol": "OXY1"}
	ret, err := DeleteMany(context.TODO(), *c, "symbol", f)
	if err != nil {
		t.Errorf("Failed to delete many documents: %v", err)
	}
	if ret == nil {
		t.Errorf("Expected to get an ID, got nil")
	} else if ret.DeletedCount == 0 {
		t.Errorf("Expected to delete one document, got %d", ret.DeletedCount)
	}
}
