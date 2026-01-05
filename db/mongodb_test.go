package db

import (
	"context"
	"testing"
	"time"

	"github.com/crowemi-io/crowemi-go-utils/config"
	"go.mongodb.org/mongo-driver/bson"
)

func setupMongoDB() *MongoClient {
	config, _ := config.Bootstrap[config.Crowemi]("../.secret/config-db.json")
	c := MongoClient{}
	c.Connect(context.TODO(), config.DatabaseURI, config.ClientName)
	return &c
}

func TestMongoDBConnect(t *testing.T) {
	p := setupMongoDB()
	if p != nil {
		t.Logf("Connected to MongoDB successfully")
	} else {
		t.Errorf("Failed to connect to MongoDB: %v", p)
	}
}
func TestMongoDBPing(t *testing.T) {
	c := setupMongoDB()
	err := c.Ping()
	if err != nil {
		t.Errorf("Failed to ping MongoDB: %v", err)
	}
}
func TestMongoDBGetOne(t *testing.T) {
	c := setupMongoDB()
	f := []MongoFilter{{Field: "symbol", Operator: "$eq", Value: "AAPL"}}
	s := []MongoSort{{Field: "buy_price", Direction: -1}}
	type order struct {
		CreatedAt time.Time `bson:"created_at"`
		Symbol    string    `bson:"symbol"`
	}
	result, err := GetOne[order](context.TODO(), c, "orders", f, s)
	if err != nil {
		t.Errorf("Failed to get one document: %v", err)
	}

	if result.Symbol != "AAPL" {
		t.Errorf("Expected symbol AAPL, got %s", result.Symbol)
	}
}
func TestMongoDBGetMany(t *testing.T) {
	c := setupMongoDB()
	// f := []MongoFilter{{Field: "symbol", Operator: "$eq", Value: "AAPL1"}}
	s := []MongoSort{{Field: "buy_price", Direction: -1}}

	type order struct {
		CreatedAt time.Time `bson:"created_at"`
		Symbol    string    `bson:"symbol"`
	}
	result, err := GetMany[order](context.TODO(), c, "orders", nil, s)
	if err != nil {
		t.Errorf("Failed to get one document: %v", err)
	}
	// Ensure the result is not empty by checking a relevant field
	if len(*result) == 0 {
		t.Errorf("Expected to get at least one document, got %d", len(*result))
	}
}
func TestMongoDBInsertOne(t *testing.T) {
	c := setupMongoDB()
	type symbol struct {
		Symbol string `bson:"symbol"`
	}
	s := symbol{Symbol: "OXY"}
	ret, err := InsertOne(context.TODO(), c, "symbol", s)
	if err != nil {
		t.Errorf("Failed to insert one document: %v", err)
	}
	if ret == nil {
		t.Errorf("Expected to get an ID, got nil")
	}
}

func TestMongoDBInsertMany(t *testing.T) {
	c := setupMongoDB()
	type symbol struct {
		Symbol string `bson:"symbol"`
	}
	s := []symbol{
		{Symbol: "OXY1"},
		{Symbol: "AAPL1"},
	}
	ret, err := InsertMany(context.TODO(), c, "symbol", s)
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
	c := setupMongoDB()
	type symbol struct {
		Symbol    string    `bson:"symbol"`
		CreatedAt time.Time `bson:"created_at"`
	}
	s := symbol{Symbol: "OXY", CreatedAt: time.Now()}
	f := []MongoFilter{{Field: "symbol", Operator: "$eq", Value: "OXY"}}
	ret, err := UpdateOne(context.TODO(), c, "symbol", f, s)
	if err != nil {
		t.Errorf("Failed to update one document: %v", err)
	}
	if ret == nil {
		t.Errorf("Expected to get an ID, got nil")
	}
}
func TestMongoDBUpdateMany(t *testing.T) {}
func TestMongoDBDeleteOne(t *testing.T) {
	c := setupMongoDB()
	f := []MongoFilter{{Field: "symbol", Operator: "$eq", Value: "OXY1"}}
	ret, err := DeleteOne(context.TODO(), c, "symbol", f)
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
	c := setupMongoDB()
	f := []MongoFilter{{Field: "symbol", Operator: "$eq", Value: "AAPL1"}}
	ret, err := DeleteMany(context.TODO(), c, "symbol", f)
	if err != nil {
		t.Errorf("Failed to delete many documents: %v", err)
	}
	if ret == nil {
		t.Errorf("Expected to get an ID, got nil")
	} else if ret.DeletedCount == 0 {
		t.Errorf("Expected to delete one document, got %d", ret.DeletedCount)
	}
}

func TestMongoDBAggregate(t *testing.T) {
	c := setupMongoDB()
	type symbol struct {
		Symbol   string  `bson:"_id"`
		Profit   float64 `bson:"profit"`
		Notional float64 `bson:"total"`
	}
	a := []bson.D{
		{{Key: "$match", Value: bson.D{{Key: "sell_at_utc", Value: nil}}}},
		{
			{Key: "$group", Value: bson.D{
				{Key: "_id", Value: "$symbol"},
				{Key: "total", Value: bson.D{{Key: "$sum", Value: 1}}},
				{Key: "profit", Value: bson.D{{Key: "$sum", Value: "$notional"}}},
			}},
		}}
	ret, err := Aggregate[symbol](context.TODO(), c, "orders", a)
	if err != nil {
		t.Errorf("Failed to aggregate documents: %v", err)
	}
	if ret == nil {
		t.Errorf("Expected to get an ID, got nil")
	}
}
