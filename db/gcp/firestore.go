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
type Update struct {
	Path  string
	Value any
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
func GetOneByID[T any](ctx context.Context, client *firestore.Client, collection string, id string) (*T, string, error) {
	var ret T
	doc, err := client.Collection(collection).Doc(id).Get(ctx)
	if status.Code(err) == codes.NotFound {
		// object not found
		return &ret, "", nil
	}
	if err != nil {
		return nil, "", err
	}
	doc.DataTo(&ret)
	return &ret, doc.Ref.ID, err
}
func GetOne[T any](ctx context.Context, client *firestore.Client, collection string, filters []db.Filter) (*T, string, error) {
	var ret T
	query := client.Collection(collection).Query
	for _, f := range filters {
		query = query.Where(f.Field, f.Operator, f.Value)
	}
	docs, err := query.Documents(ctx).GetAll()
	if err != nil {
		return nil, "", err
	}
	if len(docs) == 0 {
		return nil, "", nil
	}
	docs[0].DataTo(&ret)
	id := docs[0].Ref.ID
	return &ret, id, err
}
func InsertOne(ctx context.Context, client *firestore.Client, collection string, obj any) (*firestore.DocumentRef, *firestore.WriteResult, error) {
	ref, res, err := client.Collection(collection).Add(ctx, obj)
	if err != nil {
		return ref, res, err
	}
	return ref, res, err
}
func UpdateOne(ctx context.Context, client *firestore.Client, collection string, id string, updates []Update) (*firestore.WriteResult, error) {
	firestoreUpdates := []firestore.Update{}
	for _, update := range updates {
		firestoreUpdates = append(firestoreUpdates, firestore.Update{
			Path:  update.Path,
			Value: update.Value,
		})
	}
	ret, err := client.Collection(collection).Doc(id).Update(ctx, firestoreUpdates)
	if err != nil {
		return ret, err
	}
	return ret, err
}
func DeleteOne(ctx context.Context, client *firestore.Client, collection string, id string) (*firestore.WriteResult, error) {
	ret, err := client.Collection(collection).Doc(id).Delete(ctx)
	if err != nil {
		return ret, err
	}
	return ret, err
}

func GetMany[T any](ctx context.Context, client *firestore.Client, collection string, filters []db.Filter) (map[string]T, error) {
	// TODO: handle OR filters
	ret := map[string]T{}
	query := client.Collection(collection).Query
	for _, f := range filters {
		query = query.Where(f.Field, f.Operator, f.Value)
	}
	docs, err := query.Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}

	// no records found
	if len(docs) == 0 {
		return nil, nil
	}

	for _, doc := range docs {
		var item T
		doc.DataTo(&item)
		ret[doc.Ref.ID] = item
	}
	return ret, nil
}
func InsertMany[T any]() {}
func UpdateMany[T any]() {}
func DeleteMany[T any]() {}
