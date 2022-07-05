package repository_mysql_test

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-seidon/local/internal/mock"
	"github.com/go-seidon/local/internal/repository"
	repository_mysql "github.com/go-seidon/local/internal/repository-mysql"
	"github.com/golang/mock/gomock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRepositoryMySQL(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repository MySQL Package")
}

var _ = Describe("File Repository", func() {
	Context("NewFileRepository function", Label("unit"), func() {
		When("client is not specified", func() {
			It("should return error", func() {
				res, err := repository_mysql.NewFileRepository(nil)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("invalid client specified")))
			})
		})

		When("all parameter is specified", func() {
			It("should return result", func() {
				c := &sql.DB{}
				res, err := repository_mysql.NewFileRepository(c)

				Expect(res).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})

	Context("DeleteFile function", Label("unit"), func() {
		var (
			ctx              context.Context
			currentTimestamp time.Time
			clock            *mock.MockClock
			dbClient         sqlmock.Sqlmock
			repo             *repository_mysql.FileRepository
			p                repository.DeleteFileParam
			findFileQuery    string
			deleteFileQuery  string
			fileRows         *sqlmock.Rows
		)

		BeforeEach(func() {
			ctx = context.Background()
			t := GinkgoT()

			currentTimestamp = time.Now()
			ctrl := gomock.NewController(t)
			clock = mock.NewMockClock(ctrl)
			clock.EXPECT().Now().Return(currentTimestamp)

			db, mock, err := sqlmock.New()
			if err != nil {
				AbortSuite("failed create db mock: " + err.Error())
			}
			dbClient = mock

			repo, _ = repository_mysql.NewFileRepository(db)
			repo.Clock = clock

			p = repository.DeleteFileParam{
				UniqueId: "mock-unique-id",
				DeleteFn: func(ctx context.Context, p repository.DeleteFnParam) error {
					return nil
				},
			}
			findFileQuery = `
				SELECT 
					unique_id, name, path,
					mimetype, extension, size,
					created_at, updated_at, deleted_at
				FROM file
				WHERE unique_id = ?
			`
			deleteFileQuery = fmt.Sprintf(
				"UPDATE file SET deleted_at = '%d' WHERE unique_id = '%s'",
				currentTimestamp.UnixMilli(),
				p.UniqueId,
			)
			fileRows = sqlmock.NewRows([]string{
				"unique_id", "name", "path",
				"mimetype", "extension", "size",
				"created_at", "updated_at", "deleted_at",
			}).AddRow(
				"mock-unique-id",
				"mock-name",
				"mock-path",
				"mock-mimetype",
				"mock-extension",
				0,
				0,
				0,
				0,
			)
		})

		When("failed start db transaction", func() {
			It("should return error", func() {
				dbClient.
					ExpectBegin().
					WillReturnError(fmt.Errorf("failed start db trx"))

				res, err := repo.DeleteFile(ctx, p)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("failed start db trx")))
			})
		})

		When("failed scan record", func() {
			It("should return error", func() {
				dbClient.ExpectBegin()
				rows := sqlmock.NewRows([]string{
					"unique_id", "name", "path",
					"mimetype", "extension", "size",
					"created_at", "updated_at", "deleted_at",
				}).AddRow(
					"mock-unique-id",
					"mock-name",
					"mock-path",
					"mock-mimetype",
					"mock-extension",
					"invalid_int_value", //should be int64
					0,
					0,
					0,
				)
				dbClient.ExpectQuery(findFileQuery).WillReturnRows(rows)
				dbClient.ExpectRollback()

				res, err := repo.DeleteFile(ctx, p)

				Expect(res).To(BeNil())
				Expect(err.Error()).To(Equal("sql: Scan error on column index 5, name \"size\": converting driver.Value type string (\"invalid_int_value\") to a int64: invalid syntax"))
			})
		})

		When("record is not found", func() {
			It("should return error", func() {
				dbClient.ExpectBegin()
				dbClient.ExpectQuery(findFileQuery).
					WillReturnError(sql.ErrNoRows)
				dbClient.ExpectRollback()

				res, err := repo.DeleteFile(ctx, p)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(repository.ErrorRecordNotFound))
			})
		})

		When("failed find file record", func() {
			It("should return error", func() {
				dbClient.ExpectBegin()
				dbClient.ExpectQuery(findFileQuery).
					WillReturnError(fmt.Errorf("db error"))
				dbClient.ExpectRollback()

				res, err := repo.DeleteFile(ctx, p)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("db error")))
			})
		})

		When("failed rollback find file trx", func() {
			It("should return error", func() {
				dbClient.ExpectBegin()
				dbClient.ExpectQuery(findFileQuery).
					WillReturnError(fmt.Errorf("db error"))
				dbClient.ExpectRollback().
					WillReturnError(fmt.Errorf("failed rollback"))

				res, err := repo.DeleteFile(ctx, p)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("failed rollback")))
			})
		})

		When("failed update file record", func() {
			It("should return error", func() {
				dbClient.ExpectBegin()
				dbClient.ExpectQuery(findFileQuery).WillReturnRows(fileRows)
				dbClient.ExpectExec(deleteFileQuery).WillReturnError(fmt.Errorf("db error"))
				dbClient.ExpectRollback()

				res, err := repo.DeleteFile(ctx, p)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("db error")))
			})
		})

		When("failed rollback update file db trx", func() {
			It("should return error", func() {
				dbClient.ExpectBegin()
				dbClient.ExpectQuery(findFileQuery).WillReturnRows(fileRows)
				dbClient.ExpectExec(deleteFileQuery).WillReturnError(fmt.Errorf("db error"))
				dbClient.ExpectRollback().WillReturnError(fmt.Errorf("rollback error"))

				res, err := repo.DeleteFile(ctx, p)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("rollback error")))
			})
		})

		When("total affected row is not 1", func() {
			It("should return error", func() {
				dbClient.ExpectBegin()
				dbClient.ExpectQuery(findFileQuery).WillReturnRows(fileRows)
				dbClient.ExpectExec(deleteFileQuery).WillReturnResult(driver.ResultNoRows)
				dbClient.ExpectRollback()

				res, err := repo.DeleteFile(ctx, p)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("record is not updated")))
			})
		})

		When("failed rollback total affected row trx", func() {
			It("should return error", func() {
				dbClient.ExpectBegin()
				dbClient.ExpectQuery(findFileQuery).WillReturnRows(fileRows)
				dbClient.ExpectExec(deleteFileQuery).WillReturnResult(driver.ResultNoRows)
				dbClient.ExpectRollback().WillReturnError(fmt.Errorf("rollback error"))

				res, err := repo.DeleteFile(ctx, p)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("rollback error")))
			})
		})

		When("failed execute delete function", func() {
			It("should return error", func() {
				dbClient.ExpectBegin()
				dbClient.ExpectQuery(findFileQuery).WillReturnRows(fileRows)
				dbClient.ExpectExec(deleteFileQuery).WillReturnResult(driver.RowsAffected(1))
				p.DeleteFn = func(ctx context.Context, p repository.DeleteFnParam) error {
					return fmt.Errorf("delete fn error")
				}
				dbClient.ExpectRollback()

				res, err := repo.DeleteFile(ctx, p)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("delete fn error")))
			})
		})

		When("failed rollback delete fn db trx", func() {
			It("should return error", func() {
				dbClient.ExpectBegin()
				dbClient.ExpectQuery(findFileQuery).WillReturnRows(fileRows)
				dbClient.ExpectExec(deleteFileQuery).WillReturnResult(driver.RowsAffected(1))
				p.DeleteFn = func(ctx context.Context, p repository.DeleteFnParam) error {
					return fmt.Errorf("delete fn error")
				}
				dbClient.ExpectRollback().WillReturnError(fmt.Errorf("rollback error"))

				res, err := repo.DeleteFile(ctx, p)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("rollback error")))
			})
		})

		When("failed commit db trx", func() {
			It("should return error", func() {
				dbClient.ExpectBegin()
				dbClient.ExpectQuery(findFileQuery).WillReturnRows(fileRows)
				dbClient.ExpectExec(deleteFileQuery).WillReturnResult(driver.RowsAffected(1))
				dbClient.ExpectCommit().WillReturnError(fmt.Errorf("commit error"))

				res, err := repo.DeleteFile(ctx, p)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("commit error")))
			})
		})

		When("success delete file", func() {
			It("should return result", func() {
				dbClient.ExpectBegin()
				dbClient.ExpectQuery(findFileQuery).WillReturnRows(fileRows)
				dbClient.ExpectExec(deleteFileQuery).WillReturnResult(driver.RowsAffected(1))
				dbClient.ExpectCommit()

				res, err := repo.DeleteFile(ctx, p)

				expectedRes := &repository.DeleteFileResult{
					DeletedAt: currentTimestamp,
				}
				Expect(res).To(Equal(expectedRes))
				Expect(err).To(BeNil())
			})
		})
	})

	Context("DeleteFile function", Label("integration"), Ordered, func() {
		var (
			ctx    context.Context
			client *sql.DB
			repo   *repository_mysql.FileRepository
			p      repository.DeleteFileParam
		)

		BeforeAll(func() {
			dbClient, err := OpenDb("")
			if err != nil {
				AbortSuite("failed open test db: " + err.Error())
			}
			client = dbClient

			err = RunDbMigration(client)
			if err != nil {
				AbortSuite("failed prepare db migration: " + err.Error())
			}

			ctx = context.Background()
			repo, _ = repository_mysql.NewFileRepository(client)
		})

		BeforeEach(func() {
			p = repository.DeleteFileParam{
				UniqueId: "mock-unique-id",
				DeleteFn: func(ctx context.Context, p repository.DeleteFnParam) error {
					return nil
				},
			}
			err := InsertDummyFile(client, InsertDummyFileParam{
				UniqueId: p.UniqueId,
			})
			if err != nil {
				AbortSuite("failed prepare seed data: " + err.Error())
			}
		})

		AfterEach(func() {
			client.Exec("TRUNCATE file")
		})

		AfterAll(func() {
			client.Close()
		})

		When("file record is not available", func() {
			It("should return error", func() {
				p.UniqueId = "unavailable-unique-id"
				res, err := repo.DeleteFile(ctx, p)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(repository.ErrorRecordNotFound))
			})
		})

		When("failed proceed callback", func() {
			It("should return error", func() {
				p.DeleteFn = func(ctx context.Context, p repository.DeleteFnParam) error {
					return fmt.Errorf("failed proceed callback")
				}
				res, err := repo.DeleteFile(ctx, p)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("failed proceed callback")))
			})
		})

		When("success delete file", func() {
			It("should return result", func() {
				res, err := repo.DeleteFile(ctx, p)

				Expect(res).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})

	})

	Context("RetrieveFile function", Label("unit"), func() {
		var (
			ctx           context.Context
			dbClient      sqlmock.Sqlmock
			repo          *repository_mysql.FileRepository
			p             repository.RetrieveFileParam
			findFileQuery string
			fileRows      *sqlmock.Rows
		)

		BeforeEach(func() {
			ctx = context.Background()

			db, mock, err := sqlmock.New()
			if err != nil {
				AbortSuite("failed create db mock: " + err.Error())
			}
			dbClient = mock

			repo, _ = repository_mysql.NewFileRepository(db)

			p = repository.RetrieveFileParam{
				UniqueId: "mock-unique-id",
			}
			fileRows = sqlmock.NewRows([]string{
				"unique_id", "name", "path",
				"mimetype", "extension", "size",
				"created_at", "updated_at", "deleted_at",
			}).AddRow(
				"mock-unique-id",
				"mock-name",
				"mock-path",
				"mock-mimetype",
				"mock-extension",
				0,
				0,
				0,
				nil,
			)
		})

		When("failed scan record", func() {
			It("should return error", func() {
				rows := sqlmock.NewRows([]string{
					"unique_id", "name", "path",
					"mimetype", "extension", "size",
					"created_at", "updated_at", "deleted_at",
				}).AddRow(
					"mock-unique-id",
					"mock-name",
					"mock-path",
					"mock-mimetype",
					"mock-extension",
					"invalid_int_value", //should be int64
					0,
					0,
					0,
				)
				dbClient.ExpectQuery(findFileQuery).WillReturnRows(rows)

				res, err := repo.RetrieveFile(ctx, p)

				Expect(res).To(BeNil())
				Expect(err.Error()).To(Equal("sql: Scan error on column index 5, name \"size\": converting driver.Value type string (\"invalid_int_value\") to a int64: invalid syntax"))
			})
		})

		When("record is not found", func() {
			It("should return error", func() {
				dbClient.ExpectQuery(findFileQuery).
					WillReturnError(sql.ErrNoRows)

				res, err := repo.RetrieveFile(ctx, p)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(repository.ErrorRecordNotFound))
			})
		})

		When("failed find file record", func() {
			It("should return error", func() {
				dbClient.ExpectQuery(findFileQuery).
					WillReturnError(fmt.Errorf("db error"))

				res, err := repo.RetrieveFile(ctx, p)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("db error")))
			})
		})

		When("success find file", func() {
			It("should return result", func() {
				dbClient.ExpectQuery(findFileQuery).
					WillReturnRows(fileRows)

				res, err := repo.RetrieveFile(ctx, p)

				eRes := &repository.RetrieveFileResult{
					UniqueId:  "mock-unique-id",
					Name:      "mock-name",
					Path:      "mock-path",
					MimeType:  "mock-mimetype",
					Extension: "mock-extension",
					DeletedAt: nil,
				}
				Expect(res).To(Equal(eRes))
				Expect(err).To(BeNil())
			})
		})
	})
})
