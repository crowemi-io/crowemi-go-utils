package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	Database    *mongo.Database
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
	mc.Database = client.Database(database)
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
func GetOne[T any](ctx context.Context, client MongoClient, collection string, filter map[string]string) (T, error) {
	c := client.Database.Collection(collection)
	result := c.FindOne(ctx, FilterDocument(filter))
	var ret T
	err := result.Decode(&ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}
func GetMany[T any](ctx context.Context, client MongoClient, collection string, filter map[string]string) ([]T, error) {
	c := client.Database.Collection(collection)
	crsr, err := c.Find(ctx, FilterDocument(filter))
	if err != nil {
		return nil, err
	}
	defer crsr.Close(ctx)

	var ret []T
	err = crsr.All(ctx, &ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}
func InsertOne[T any](ctx context.Context, client MongoClient, collection string, doc T) (*mongo.InsertOneResult, error) {
	c := client.Database.Collection(collection)
	ret, err := c.InsertOne(ctx, doc)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
func InsertMany[T any](ctx context.Context, client MongoClient, collection string, docs []T) (*mongo.InsertManyResult, error) {
	c := client.Database.Collection(collection)
	IDocs := make([]interface{}, len(docs))
	for i, doc := range docs {
		IDocs[i] = doc
	}
	ret, err := c.InsertMany(ctx, IDocs)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
func UpdateOne[T any](ctx context.Context, client MongoClient, collection string, filter map[string]string, doc T) (*mongo.UpdateResult, error) {
	c := client.Database.Collection(collection)
	ret, err := c.UpdateOne(ctx, FilterDocument(filter), bson.M{"$set": doc})
	if err != nil {
		return nil, err
	}
	return ret, nil
}
func UpdateMany[T any](ctx context.Context, client MongoClient, collection string, filter map[string]string, doc []T) (*mongo.UpdateResult, error) {
	c := client.Database.Collection(collection)
	ret, err := c.UpdateMany(ctx, FilterDocument(filter), doc)
	if err != nil {
		return ret, err
	}
	return ret, nil
}
func DeleteOne(ctx context.Context, client MongoClient, collection string, filter map[string]string) (*mongo.DeleteResult, error) {
	c := client.Database.Collection(collection)
	ret, err := c.DeleteOne(ctx, FilterDocument(filter))
	if err != nil {
		return nil, err
	}
	return ret, nil
}
func DeleteMany(ctx context.Context, client MongoClient, collection string, filter map[string]string) (*mongo.DeleteResult, error) {
	c := client.Database.Collection(collection)
	ret, err := c.DeleteMany(ctx, FilterDocument(filter))
	if err != nil {
		return nil, err
	}
	return ret, nil
}
