package db

type IDatabase interface {
	GetOne()
	GetMany()
	InsertOne()
	InsertMany()
	UpdateOne()
	UpdateMany()
	DeleteOne()
	DeleteMany()
	Connect(uri string) error
}
