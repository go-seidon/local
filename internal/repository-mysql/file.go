package repository_mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/go-seidon/local/internal/repository"
)

type FileRepository struct {
	client *sql.DB
}

func (r *FileRepository) DeleteFile(ctx context.Context, p repository.DeleteFileParam, o repository.DeleteFileOpt) (*repository.DeleteFileResult, error) {
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
		txErr := tx.Rollback()
		if txErr != nil {
			return nil, err
		}
		return nil, err
	}

	sqlQuery := fmt.Sprintf(
		"UPDATE file SET deleted_at = '%s' WHERE unique_id = '%s'",
		currentTimestamp.Format("2006-01-02 15:04:05"),
		file.UniqueId,
	)
	qRes, err := tx.Exec(sqlQuery)
	if err != nil {
		txErr := tx.Rollback()
		if txErr != nil {
			return nil, err
		}
		return nil, err
	}

	// error is ommited since mysql driver is able to returning totalAffected
	totalAffected, _ := qRes.RowsAffected()

	if totalAffected != 1 {
		txErr := tx.Rollback()
		if txErr != nil {
			return nil, err
		}
		return nil, fmt.Errorf("record is not updated")
	}

	err = o.DeleteFn(ctx, repository.DeleteFnParam{
		FilePath: file.Path,
	})
	if err != nil {
		txErr := tx.Rollback()
		if txErr != nil {
			return nil, err
		}
		return nil, err
	}

	txErr := tx.Commit()
	if txErr != nil {
		return nil, err
	}

	res := &repository.DeleteFileResult{
		DeletedAt: currentTimestamp,
	}
	return res, nil
}

func (r *FileRepository) findFile(ctx context.Context, p findFileParam) (*findFileResult, error) {
	var client Client
	client = r.client

	if p.DbTransaction != nil {
		client = p.DbTransaction
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
	err := client.QueryRow(sqlQuery, p.UniqueId).Scan(
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
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func NewFileRepository(client *sql.DB) (*FileRepository, error) {
	if client == nil {
		return nil, fmt.Errorf("invalid client specified")
	}

	r := &FileRepository{
		client: client,
	}
	return r, nil
}
