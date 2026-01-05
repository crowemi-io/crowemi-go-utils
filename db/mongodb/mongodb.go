package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	Database    *mongo.Database
	MongoClient *mongo.Client
}
type MongoFilter struct {
	Field    string
	Operator string
	Value    interface{}
}
type MongoSort struct {
	Field     string
	Direction int
}
type MongoAggregate struct {
	Field    string
	Operator string
	Value    interface{}
}

func (mc *MongoClient) Ping() error {
	return mc.MongoClient.Ping(context.TODO(), nil)
}
func (mc *MongoClient) Connect(ctx context.Context, uri string, database string) error {
	// pull parameters from config
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

func createFilter(filter []MongoFilter) (primitive.M, error) {
	var f = bson.M{}
	for _, value := range filter {
		switch value.Operator {
		case "$eq", "$ne", "$gt", "$gte", "$lt", "$lte":
			// Simple operators: { field: { $op: value } }
			f[value.Field] = bson.M{value.Operator: value.Value}
		case "$in", "$nin":
			// Array operators: { field: { $in: []interface{} } }
			if values, ok := value.Value.([]interface{}); ok {
				f[value.Field] = bson.M{value.Operator: values}
			} else {
				return nil, fmt.Errorf("value for %s must be a slice", value.Operator)
			}
		default:
			return nil, fmt.Errorf("unsupported operator: %s", value.Operator)
		}
	}
	return f, nil
}
func createSort(sort []MongoSort) primitive.D {
	var ret = bson.D{}
	for _, value := range sort {
		ret = append(ret, bson.E{value.Field, value.Direction})
	}
	return ret
}

func GetOne[T any](ctx context.Context, client *MongoClient, collection string, filter []MongoFilter, sort []MongoSort) (T, error) {
	var ret T
	var options options.FindOneOptions
	var err error

	c := client.Database.Collection(collection)
	f := primitive.M{}
	if filter != nil {
		f, err = createFilter(filter)
		if err != nil {
			return ret, err
		}
	}
	if sort != nil {
		sortOptions := createSort(sort)
		options.SetSort(sortOptions)
		if err != nil {
			return ret, err
		}
	}
	result := c.FindOne(ctx, f, &options)
	err = result.Decode(&ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}
func GetMany[T any](ctx context.Context, client *MongoClient, collection string, filter []MongoFilter, sort []MongoSort) (*[]T, error) {
	var ret []T
	var err error
	var options options.FindOptions

	c := client.Database.Collection(collection)
	f := primitive.M{}

	if filter != nil {
		f, err = createFilter(filter)
		if err != nil {
			return nil, fmt.Errorf("GetMany: failed to create filter: %w", err)
		}
	}
	if sort != nil {
		sortOptions := createSort(sort)
		options.SetSort(sortOptions)
		if err != nil {
			return &ret, err
		}
	}

	crsr, err := c.Find(ctx, f, &options)
	if err != nil {
		return nil, fmt.Errorf("GetMany: failed to find: %w", err)
	}
	defer crsr.Close(ctx)

	err = crsr.All(ctx, &ret)
	if err != nil {
		return &ret, fmt.Errorf("GetMany: failed to decode: %w", err)
	}
	return &ret, nil
}
func InsertOne[T any](ctx context.Context, client *MongoClient, collection string, doc T) (*mongo.InsertOneResult, error) {
	c := client.Database.Collection(collection)
	ret, err := c.InsertOne(ctx, doc)
	if err != nil {
		return nil, fmt.Errorf("InsertOne: failed to insert one: %w", err)
	}
	return ret, nil
}
func InsertMany[T any](ctx context.Context, client *MongoClient, collection string, docs []T) (*mongo.InsertManyResult, error) {
	c := client.Database.Collection(collection)
	IDocs := make([]interface{}, len(docs))
	for i, doc := range docs {
		IDocs[i] = doc
	}
	ret, err := c.InsertMany(ctx, IDocs)
	if err != nil {
		return nil, fmt.Errorf("InsertMany: failed to insert many: %w", err)
	}
	return ret, nil
}
func UpdateOne[T any](ctx context.Context, client *MongoClient, collection string, filter []MongoFilter, doc T) (*mongo.UpdateResult, error) {
	c := client.Database.Collection(collection)
	f, err := createFilter(filter)
	if err != nil {
		return nil, fmt.Errorf("UpdateOne: failed to create filter: %w", err)
	}

	ret, err := c.UpdateOne(ctx, f, bson.M{"$set": doc})
	if err != nil {
		return nil, fmt.Errorf("UpdateOne: failed to update one: %w", err)
	}
	return ret, nil
}
func UpdateMany[T any](ctx context.Context, client *MongoClient, collection string, filter []MongoFilter, doc []T) (*mongo.UpdateResult, error) {
	c := client.Database.Collection(collection)
	f, err := createFilter(filter)
	if err != nil {
		return nil, fmt.Errorf("UpdateMany: failed to create filter: %w", err)
	}
	ret, err := c.UpdateMany(ctx, f, doc)
	if err != nil {
		return ret, fmt.Errorf("UpdateMany: failed to update many: %w", err)
	}
	return ret, nil
}
func DeleteOne(ctx context.Context, client *MongoClient, collection string, filter []MongoFilter) (*mongo.DeleteResult, error) {
	c := client.Database.Collection(collection)
	f, err := createFilter(filter)
	if err != nil {
		return nil, fmt.Errorf("DeleteOne: failed to create filter: %w", err)
	}
	ret, err := c.DeleteOne(ctx, f)
	if err != nil {
		return nil, fmt.Errorf("DeleteOne: failed to delete one: %w", err)
	}
	return ret, nil
}
func DeleteMany(ctx context.Context, client *MongoClient, collection string, filter []MongoFilter) (*mongo.DeleteResult, error) {
	c := client.Database.Collection(collection)
	f, err := createFilter(filter)
	if err != nil {
		return nil, fmt.Errorf("DeleteMany: failed to create filter: %w", err)
	}
	ret, err := c.DeleteMany(ctx, f)
	if err != nil {
		return nil, fmt.Errorf("DeleteMany: failed to delete many: %w", err)
	}
	return ret, nil
}
func Aggregate[T any](ctx context.Context, client *MongoClient, collection string, pipeline []bson.D) ([]T, error) {
	c := client.Database.Collection(collection)
	crsr, err := c.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("Aggregate: failed to aggregate: %w", err)
	}
	defer crsr.Close(ctx)

	var ret []T
	err = crsr.All(ctx, &ret)
	if err != nil {
		return ret, fmt.Errorf("Aggregate: failed to decode: %w", err)
	}
	return ret, nil
}
