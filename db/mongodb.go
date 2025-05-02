package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	Database    string
	MongoClient *mongo.Client
}

func (mc *MongoClient) Ping() error {
	return mc.MongoClient.Ping(context.TODO(), nil)
}
func (mc *MongoClient) Connect(ctx context.Context, uri string, database string) error {
	options := options.Client().ApplyURI(uri).SetMaxPoolSize(10).SetMinPoolSize(1).SetMaxConnIdleTime(10)
	client, err := mongo.Connect(ctx, options)
	if err != nil {
		return err
	}
	mc.MongoClient = client
	mc.Database = database
	return nil
}
func (mc *MongoClient) Disconnect() error {
	return mc.MongoClient.Disconnect(context.TODO())
}

func FilterDocument(filter map[string]string) bson.D {
	var bsonFilters = bson.D{}
	for key, value := range filter {
		bsonFilters = append(bsonFilters, bson.E{Key: key, Value: value})
	}
	return bsonFilters
}
func GetOne[T any](client MongoClient, collection string, filter map[string]string) (T, error) {
	c := client.MongoClient.Database(client.Database).Collection(collection)
	result := c.FindOne(context.TODO(), FilterDocument(filter))
	var ret T
	err := result.Decode(&ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}
func GetMany[T any](client MongoClient, collection string, filter map[string]string) ([]T, error) {
	c := client.MongoClient.Database(client.Database).Collection(collection)
	crsr, err := c.Find(context.TODO(), FilterDocument(filter))
	if err != nil {
		return nil, err
	}

	var ret []T
	err = crsr.All(context.TODO(), &ret)
	if err != nil {
		return ret, err
	}
	crsr.Close(context.TODO())
	return ret, nil
}
func InsertOne[T any]()  {}
func InsertMany[T any]() {}
func UpdateOne[T any]()  {}
func UpdateMany[T any]() {}
func DeleteOne[T any]()  {}
func DeleteMany[T any]() {}
