package repository_mysql_test

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-seidon/local/internal/mock"
	"github.com/go-seidon/local/internal/repository"
	repository_mysql "github.com/go-seidon/local/internal/repository-mysql"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("OAuth Repository", func() {

	Context("NewOAuthRepository function", Label("unit"), func() {
		When("db client is not specified", func() {
			It("should return error", func() {
				res, err := repository_mysql.NewOAuthRepository()

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("invalid db client specified")))
			})
		})

		When("required parameter is specified", func() {
			It("should return result", func() {
				opt := repository_mysql.WithDbClient(&sql.DB{})
				res, err := repository_mysql.NewOAuthRepository(opt)

				Expect(res).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})

		When("clock is specified", func() {
			It("should return result", func() {
				clockOpt := repository_mysql.WithClock(&mock.MockClock{})
				dbOpt := repository_mysql.WithDbClient(&sql.DB{})
				res, err := repository_mysql.NewOAuthRepository(clockOpt, dbOpt)

				Expect(res).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})

	Context("FindClient function", Label("unit"), func() {
		var (
			ctx             context.Context
			dbClient        sqlmock.Sqlmock
			repo            repository.OAuthRepository
			p               repository.FindClientParam
			findClientQuery string
		)

		BeforeEach(func() {
			ctx = context.Background()

			db, mock, err := sqlmock.New()
			if err != nil {
				AbortSuite("failed create db mock: " + err.Error())
			}
			dbClient = mock

			dbOpt := repository_mysql.WithDbClient(db)
			repo, _ = repository_mysql.NewOAuthRepository(dbOpt)
			p = repository.FindClientParam{
				ClientId: "client_id",
			}

			findClientQuery = regexp.QuoteMeta(`
				SELECT 
					client_id, client_secret
				FROM oauth_client
				WHERE client_id = ?
			`)
		})

		When("failed client not found", func() {
			It("should return error", func() {
				dbClient.
					ExpectQuery(findClientQuery).
					WillReturnError(sql.ErrNoRows)

				res, err := repo.FindClient(ctx, p)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(repository.ErrorRecordNotFound))
			})
		})

		When("unexpected error happened", func() {
			It("should return error", func() {
				dbClient.
					ExpectQuery(findClientQuery).
					WillReturnError(fmt.Errorf("error"))

				res, err := repo.FindClient(ctx, p)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("error")))
			})
		})

		When("client is available", func() {
			It("should return result", func() {
				rows := sqlmock.NewRows([]string{
					"client_id", "client_secret",
				}).AddRow(
					"mock-client-id",
					"mock-client-client_secret",
				)
				dbClient.
					ExpectQuery(findClientQuery).
					WillReturnRows(rows)

				res, err := repo.FindClient(ctx, p)

				expectedRes := &repository.FindClientResult{
					ClientId:     "mock-client-id",
					ClientSecret: "mock-client-client_secret",
				}
				Expect(res).To(Equal(expectedRes))
				Expect(err).To(BeNil())
			})
		})
	})

})
