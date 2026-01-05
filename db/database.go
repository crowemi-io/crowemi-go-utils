package db

import "context"

type IDatabase interface {
	Connect(ctx context.Context) error
}

type Filter struct {
	Field    string
	Operator string
	Value    interface{}
}
type Sort struct {
	Field     string
	Direction int
}
