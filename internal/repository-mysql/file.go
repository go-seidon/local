package repository_mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/go-seidon/local/internal/repository"
)

type fileRepository struct {
	client *sql.DB
}

func (r *fileRepository) DeleteFile(ctx context.Context, p repository.DeleteFileParam, o repository.DeleteFileOpt) (*repository.DeleteFileResult, error) {
	currentTimestamp := time.Now()

	tx, err := r.client.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
	})
	if err != nil {
		return nil, err
	}

	file, err := r.findFile(ctx, findFileParam{
		UniqueId:      p.UniqueId,
		DbTransaction: tx,
		ShouldLock:    true,
	})
	if err != nil {
		return nil, err
	}

	sqlQuery := `
		UPDATE file 
		SET deleted_at = $2 
		WHERE unique_id = $1
	`
	qRes, err := tx.Exec(sqlQuery, file.UniqueId, currentTimestamp)
	if err != nil {
		return nil, err
	}

	totalAffected, err := qRes.RowsAffected()
	if err != nil {
		return nil, err
	}

	if totalAffected != 1 {
		return nil, fmt.Errorf("record is not updated")
	}

	err = o.DeleteFn(ctx, repository.DeleteFnParam{
		FilePath: file.Path,
	})
	if err != nil {
		return nil, err
	}

	res := &repository.DeleteFileResult{
		DeletedAt: currentTimestamp,
	}
	return res, nil
}

func (r *fileRepository) findFile(ctx context.Context, p findFileParam) (*findFileResult, error) {
	var client Client
	client = r.client

	if p.DbTransaction != nil {
		client = p.DbTransaction
	}

	sqlQuery := `
		SELECT 
			unique_id, name, path,
			mimetype, extension, size
			created_at, updated_at, deleted_at
		FROM file 
	`
	if p.ShouldLock {
		sqlQuery += ` LOCK FOR UPDATE `
	}
	sqlQuery += ` WHERE unique_id = $1 `

	row := client.QueryRow(sqlQuery, p.UniqueId)
	err := row.Err()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrorRecordNotFound
		}
		return nil, err
	}

	var res *findFileResult
	err = row.Scan(
		&res.UniqueId,
		&res.Name,
		&res.Path,
		&res.MimeType,
		&res.Extension,
		&res.Size,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.DeletedAt,
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type findFileParam struct {
	UniqueId      string
	ShouldLock    bool
	DbTransaction *sql.Tx
}

type findFileResult struct {
	UniqueId  string
	Name      string
	Path      string
	MimeType  string
	Extension string
	Size      int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func NewFileRepository(client *sql.DB) (*fileRepository, error) {
	if client == nil {
		return nil, fmt.Errorf("invalid repo client specified")
	}

	r := &fileRepository{
		client: client,
	}
	return r, nil
}
