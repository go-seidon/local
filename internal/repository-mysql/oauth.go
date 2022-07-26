package repository_mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-seidon/local/internal/datetime"
	"github.com/go-seidon/local/internal/repository"
)

type oAuthRepository struct {
	dbClient *sql.DB
	clock    datetime.Clock
}

func (r *oAuthRepository) FindClient(ctx context.Context, p repository.FindClientParam) (*repository.FindClientResult, error) {
	sqlQuery := `
		SELECT 
			client_id, client_secret
		FROM oauth_client
		WHERE client_id = ?
	`

	var res repository.FindClientResult
	row := r.dbClient.QueryRow(sqlQuery, p.ClientId)
	err := row.Scan(
		&res.ClientId,
		&res.ClientSecret,
	)
	if err == nil {
		return &res, nil
	}

	if errors.Is(err, sql.ErrNoRows) {
		return nil, repository.ErrorRecordNotFound
	}
	return nil, err
}

func NewOAuthRepository(opts ...RepoOption) (*oAuthRepository, error) {
	option := RepositoryOption{}
	for _, opt := range opts {
		opt(&option)
	}

	if option.dbClient == nil {
		return nil, fmt.Errorf("invalid db client specified")
	}

	var clock datetime.Clock
	if option.clock == nil {
		clock = datetime.NewClock()
	} else {
		clock = option.clock
	}

	r := &oAuthRepository{
		dbClient: option.dbClient,
		clock:    clock,
	}
	return r, nil
}
