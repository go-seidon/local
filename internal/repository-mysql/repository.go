package repository_mysql

import (
	"database/sql"

	"github.com/go-seidon/local/internal/datetime"
)

type RepositoryOption struct {
	dbClient *sql.DB
	clock    datetime.Clock
}

type RepoOption = func(*RepositoryOption)

func WithDbClient(dbClient *sql.DB) RepoOption {
	return func(ro *RepositoryOption) {
		ro.dbClient = dbClient
	}
}

func WithClock(clock datetime.Clock) RepoOption {
	return func(ro *RepositoryOption) {
		ro.clock = clock
	}
}
