package repository_mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-seidon/local/internal/datetime"
	"github.com/go-seidon/local/internal/repository"
)

type FileRepository struct {
	client *sql.DB
	Clock  datetime.Clock
}

func (r *FileRepository) DeleteFile(ctx context.Context, p repository.DeleteFileParam) (*repository.DeleteFileResult, error) {
	currentTimestamp := r.Clock.Now()

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
		txErr := tx.Rollback()
		if txErr != nil {
			return nil, txErr
		}
		return nil, err
	}

	sqlQuery := fmt.Sprintf(
		"UPDATE file SET deleted_at = '%d' WHERE unique_id = '%s'",
		currentTimestamp.UnixMilli(),
		file.UniqueId,
	)
	qRes, err := tx.Exec(sqlQuery)
	if err != nil {
		txErr := tx.Rollback()
		if txErr != nil {
			return nil, txErr
		}
		return nil, err
	}

	// error is ommited since mysql driver is able to returning totalAffected
	totalAffected, _ := qRes.RowsAffected()
	if totalAffected != 1 {
		txErr := tx.Rollback()
		if txErr != nil {
			return nil, txErr
		}
		return nil, fmt.Errorf("record is not updated")
	}

	err = p.DeleteFn(ctx, repository.DeleteFnParam{
		FilePath: file.Path,
	})
	if err != nil {
		txErr := tx.Rollback()
		if txErr != nil {
			return nil, txErr
		}
		return nil, err
	}

	txErr := tx.Commit()
	if txErr != nil {
		return nil, txErr
	}

	res := &repository.DeleteFileResult{
		DeletedAt: currentTimestamp,
	}
	return res, nil
}

func (r *FileRepository) RetrieveFile(ctx context.Context, p repository.RetrieveFileParam) (*repository.RetrieveFileResult, error) {
	file, err := r.findFile(ctx, findFileParam{
		UniqueId: p.UniqueId,
	})
	if err != nil {
		return nil, err
	}
	res := &repository.RetrieveFileResult{
		UniqueId:  file.UniqueId,
		Name:      file.Name,
		Path:      file.Path,
		MimeType:  file.MimeType,
		Extension: file.Extension,
		DeletedAt: file.DeletedAt,
	}
	return res, nil
}

func (r *FileRepository) findFile(ctx context.Context, p findFileParam) (*findFileResult, error) {
	var q Query
	q = r.client

	if p.DbTransaction != nil {
		q = p.DbTransaction
	}

	sqlQuery := `
		SELECT 
			unique_id, name, path,
			mimetype, extension, size,
			created_at, updated_at, deleted_at
		FROM file
		WHERE unique_id = ?
	`
	if p.ShouldLock {
		sqlQuery += ` FOR UPDATE `
	}

	var res findFileResult
	row := q.QueryRow(sqlQuery, p.UniqueId)
	err := row.Scan(
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
	if err == nil {
		return &res, nil
	}

	if errors.Is(err, sql.ErrNoRows) {
		return nil, repository.ErrorRecordNotFound
	}
	return nil, err
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
	CreatedAt int64
	UpdatedAt int64
	DeletedAt *int64
}

func NewFileRepository(client *sql.DB) (*FileRepository, error) {
	if client == nil {
		return nil, fmt.Errorf("invalid client specified")
	}

	clock := datetime.NewClock()
	r := &FileRepository{
		client: client,
		Clock:  clock,
	}
	return r, nil
}
