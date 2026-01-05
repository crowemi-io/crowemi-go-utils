package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/crowemi-io/crowemi-go-utils/config"
	"github.com/crowemi-io/crowemi-go-utils/db"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Client struct {
	Config *config.GoogleCloud
}
type ConnectOptions struct {
	ProjectID string
	Database  string
}

func (fc *Client) Connect(ctx context.Context, options ...ConnectOptions) (*firestore.Client, error) {
	projectID := fc.Config.ProjectID
	database := fc.Config.Firestore.Database

	// TODO: add implement options override
	client, err := firestore.NewClientWithDatabase(ctx, projectID, database)
	if err != nil {
		return nil, err
	}
	return client, err
}
func GetOne[T any](ctx context.Context, client *firestore.Client, collection string, id string) (*T, error) {
	var ret T
	doc, err := client.Collection(collection).Doc(id).Get(ctx)
	if status.Code(err) == codes.NotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	doc.DataTo(&ret)
	return &ret, err
}
func InsertOne[T any](ctx context.Context, client *firestore.Client, collection string, obj T) (*firestore.DocumentRef, *firestore.WriteResult, error) {
	ref, res, err := client.Collection(collection).Add(ctx, obj)
	if err != nil {
		return ref, res, err
	}
	return ref, res, err
}
func UpdateOne[T any]() {}

func GetMany[T any](ctx context.Context, client *firestore.Client, collection string, filters []db.Filter) (*[]T, error) {
	// TODO: handle OR filters
	var ret []T
	query := client.Collection(collection).Query
	for _, f := range filters {
		query = query.Where(f.Field, f.Operator, f.Value)
	}
	docs, err := query.Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	for _, doc := range docs {
		var item T
		doc.DataTo(&item)
		ret = append(ret, item)
	}
	return &ret, nil
}
func InsertMany[T any]() {}
func UpdateMany[T any]() {}
func DeleteOne[T any]()  {}
func DeleteMany[T any]() {}
