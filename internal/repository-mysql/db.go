package repository_mysql

import (
	"context"
	"database/sql"
)

type Client interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	Query
}

type Transaction interface {
	Commit() error
	Rollback() error
	Query
}

type Query interface {
	QueryRow(query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
}
