package repository

import "context"

type Connection interface {
	Transaction
}

type Transaction interface {
	Start(ctx context.Context) error
	Rollback(ctx context.Context) error
	Commit(ctx context.Context) error
}
