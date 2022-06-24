package repository_mysql_test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/go-seidon/local/internal/repository"
	repository_mysql "github.com/go-seidon/local/internal/repository-mysql"
	"github.com/go-sql-driver/mysql"

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

	Context("DeleteFile function", Label("integration"), Ordered, func() {
		var (
			ctx    context.Context
			client *sql.DB
			repo   *repository_mysql.FileRepository
			p      repository.DeleteFileParam
			o      repository.DeleteFileOpt
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
			}
			o = repository.DeleteFileOpt{
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

		When("failed start db transaction", func() {
			It("should return error", func() {
				invalidClient, _ := OpenDb("invalid_username:invalid_password@tcp(localhost:3307)/invalid_database")
				invalidRepo, _ := repository_mysql.NewFileRepository(invalidClient)
				res, err := invalidRepo.DeleteFile(ctx, p, o)

				expectedErr := &mysql.MySQLError{
					Number:  1045,
					Message: "Access denied for user 'invalid_username'@'172.31.0.1' (using password: YES)",
				}
				Expect(res).To(BeNil())
				Expect(err).To(Equal(expectedErr))
			})
		})

		When("file record is not available", func() {
			It("should return error", func() {
				p.UniqueId = "unavailable-unique-id"
				res, err := repo.DeleteFile(ctx, p, o)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(repository.ErrorRecordNotFound))
			})
		})

		When("failed proceed callback", func() {
			It("should return error", func() {
				o.DeleteFn = func(ctx context.Context, p repository.DeleteFnParam) error {
					return fmt.Errorf("failed proceed callback")
				}
				res, err := repo.DeleteFile(ctx, p, o)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("failed proceed callback")))
			})
		})

		When("success delete file", func() {
			It("should return result", func() {
				res, err := repo.DeleteFile(ctx, p, o)

				Expect(res).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})

	})
})
