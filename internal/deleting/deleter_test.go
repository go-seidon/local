package deleting_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-seidon/local/internal/deleting"
	"github.com/go-seidon/local/internal/explorer"
	"github.com/go-seidon/local/internal/mock"
	"github.com/go-seidon/local/internal/repository"
	"github.com/golang/mock/gomock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDeleting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Deleting Package")
}

var _ = Describe("Deleter Service", func() {
	Context("NewDeleter function", func() {
		var (
			fileRepo    *mock.MockFileRepository
			fileManager *mock.MockFileManager
			logger      *mock.MockLogger
			p           deleting.NewDeleterParam
		)

		BeforeEach(func() {
			t := GinkgoT()
			ctrl := gomock.NewController(t)
			fileRepo = mock.NewMockFileRepository(ctrl)
			fileManager = mock.NewMockFileManager(ctrl)
			logger = mock.NewMockLogger(ctrl)
			p = deleting.NewDeleterParam{
				FileRepo:    fileRepo,
				FileManager: fileManager,
				Logger:      logger,
			}
		})

		When("success create service", func() {
			It("should return result", func() {
				res, err := deleting.NewDeleter(p)

				Expect(res).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})

		When("file repo is not specified", func() {
			It("should return error", func() {
				p.FileRepo = nil
				res, err := deleting.NewDeleter(p)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("file repo is not specified")))
			})
		})

		When("file manager is not specified", func() {
			It("should return error", func() {
				p.FileManager = nil
				res, err := deleting.NewDeleter(p)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("file manager is not specified")))
			})
		})

		When("logger is not specified", func() {
			It("should return error", func() {
				p.Logger = nil
				res, err := deleting.NewDeleter(p)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("logger is not specified")))
			})
		})
	})

	Context("DeleteFile function", func() {
		var (
			ctx         context.Context
			p           deleting.DeleteFileParam
			fileRepo    *mock.MockFileRepository
			fileManager *mock.MockFileManager
			log         *mock.MockLogger
			dbConn      *mock.MockConnection
			s           deleting.Deleter
			findParam   repository.FindFileParam
			findRes     *repository.FindFileResult
			deleteParam repository.DeleteFileParam
			deleteRes   *repository.DeleteFileResult
			removeParam explorer.RemoveFileParam
			removeRes   *explorer.RemoveFileResult
			finalRes    *deleting.DeleteFileResult
		)

		BeforeEach(func() {
			currentTimestamp := time.Now()
			ctx = context.Background()
			p = deleting.DeleteFileParam{
				FileId: "mock-file-id",
			}
			t := GinkgoT()
			ctrl := gomock.NewController(t)
			fileRepo = mock.NewMockFileRepository(ctrl)
			fileManager = mock.NewMockFileManager(ctrl)
			log = mock.NewMockLogger(ctrl)
			dbConn = mock.NewMockConnection(ctrl)
			s, _ = deleting.NewDeleter(deleting.NewDeleterParam{
				FileRepo:    fileRepo,
				FileManager: fileManager,
				Logger:      log,
			})
			findParam = repository.FindFileParam{
				UniqueId:     p.FileId,
				DbConnection: dbConn,
			}
			findRes = &repository.FindFileResult{
				UniqueId: "mock-unique-id",
				Name:     "mock-file-name",
				Path:     "mock/path",
			}
			deleteParam = repository.DeleteFileParam{
				UniqueId:     findRes.UniqueId,
				DbConnection: dbConn,
			}
			deleteRes = &repository.DeleteFileResult{
				DeletedAt: currentTimestamp,
			}
			removeParam = explorer.RemoveFileParam{
				Path: findRes.Path,
			}
			removeRes = &explorer.RemoveFileResult{
				RemovedAt: currentTimestamp,
			}
			finalRes = &deleting.DeleteFileResult{
				DeletedAt: currentTimestamp,
			}

			log.EXPECT().
				Debug("In function: DeleteFile").
				Times(1)
			log.EXPECT().
				Debug("Returning function: DeleteFile").
				Times(1)
		})

		When("file id is not specified", func() {
			It("should return error", func() {
				p.FileId = ""
				res, err := s.DeleteFile(ctx, p)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("invalid file id parameter")))
			})
		})

		When("failed start db transaction", func() {
			It("should return error", func() {
				fileRepo.
					EXPECT().
					GetConnection().
					Return(dbConn).
					Times(1)
				dbConn.
					EXPECT().
					Start(gomock.Eq(ctx)).
					Return(fmt.Errorf("network error")).
					Times(1)

				res, err := s.DeleteFile(ctx, p)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("network error")))
			})
		})

		When("failed rollback find file query", func() {
			It("should return error", func() {
				fileRepo.
					EXPECT().
					GetConnection().
					Return(dbConn).
					Times(1)
				dbConn.
					EXPECT().
					Start(gomock.Eq(ctx)).
					Return(nil).
					Times(1)

				fileRepo.
					EXPECT().
					FindFile(gomock.Eq(ctx), gomock.Eq(findParam)).
					Return(nil, fmt.Errorf("some error")).
					Times(1)

				dbConn.
					EXPECT().
					Rollback(gomock.Eq(ctx)).
					Return(fmt.Errorf("rollback error")).
					Times(1)

				res, err := s.DeleteFile(ctx, p)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("rollback error")))
			})
		})

		When("file is not available in database", func() {
			It("should return error", func() {
				fileRepo.
					EXPECT().
					GetConnection().
					Return(dbConn).
					Times(1)
				dbConn.
					EXPECT().
					Start(gomock.Eq(ctx)).
					Return(nil).
					Times(1)

				fileRepo.
					EXPECT().
					FindFile(gomock.Eq(ctx), gomock.Eq(findParam)).
					Return(nil, repository.ErrorRecordNotFound).
					Times(1)

				dbConn.
					EXPECT().
					Rollback(gomock.Eq(ctx)).
					Return(nil).
					Times(1)

				res, err := s.DeleteFile(ctx, p)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(deleting.ErrorResourceNotFound))
			})
		})

		When("failed find file in database", func() {
			It("should return error", func() {
				fileRepo.
					EXPECT().
					GetConnection().
					Return(dbConn).
					Times(1)
				dbConn.
					EXPECT().
					Start(gomock.Eq(ctx)).
					Return(nil).
					Times(1)

				fileRepo.
					EXPECT().
					FindFile(gomock.Eq(ctx), gomock.Eq(findParam)).
					Return(nil, fmt.Errorf("network error")).
					Times(1)

				dbConn.
					EXPECT().
					Rollback(gomock.Eq(ctx)).
					Return(nil).
					Times(1)

				res, err := s.DeleteFile(ctx, p)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("network error")))
			})
		})

		When("failed rollback delete file query", func() {
			It("should return error", func() {
				fileRepo.
					EXPECT().
					GetConnection().
					Return(dbConn).
					Times(1)
				dbConn.
					EXPECT().
					Start(gomock.Eq(ctx)).
					Return(nil).
					Times(1)

				fileRepo.
					EXPECT().
					FindFile(gomock.Eq(ctx), gomock.Eq(findParam)).
					Return(findRes, nil).
					Times(1)

				fileRepo.
					EXPECT().
					DeleteFile(gomock.Eq(ctx), gomock.Eq(deleteParam)).
					Return(nil, fmt.Errorf("db error")).
					Times(1)

				dbConn.
					EXPECT().
					Rollback(gomock.Eq(ctx)).
					Return(fmt.Errorf("rollback error")).
					Times(1)

				res, err := s.DeleteFile(ctx, p)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("rollback error")))
			})
		})

		When("failed delete file in database", func() {
			It("should return error", func() {
				fileRepo.
					EXPECT().
					GetConnection().
					Return(dbConn).
					Times(1)
				dbConn.
					EXPECT().
					Start(gomock.Eq(ctx)).
					Return(nil).
					Times(1)

				fileRepo.
					EXPECT().
					FindFile(gomock.Eq(ctx), gomock.Eq(findParam)).
					Return(findRes, nil).
					Times(1)

				fileRepo.
					EXPECT().
					DeleteFile(gomock.Eq(ctx), gomock.Eq(deleteParam)).
					Return(nil, fmt.Errorf("db error")).
					Times(1)

				dbConn.
					EXPECT().
					Rollback(gomock.Eq(ctx)).
					Return(nil).
					Times(1)

				res, err := s.DeleteFile(ctx, p)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("db error")))
			})
		})

		When("failed rollback remove disk file", func() {
			It("should return error", func() {
				fileRepo.
					EXPECT().
					GetConnection().
					Return(dbConn).
					Times(1)
				dbConn.
					EXPECT().
					Start(gomock.Eq(ctx)).
					Return(nil).
					Times(1)

				fileRepo.
					EXPECT().
					FindFile(gomock.Eq(ctx), gomock.Eq(findParam)).
					Return(findRes, nil).
					Times(1)

				fileRepo.
					EXPECT().
					DeleteFile(gomock.Eq(ctx), gomock.Eq(deleteParam)).
					Return(deleteRes, nil).
					Times(1)

				fileManager.
					EXPECT().
					RemoveFile(gomock.Eq(ctx), gomock.Eq(removeParam)).
					Return(nil, fmt.Errorf("r/w error")).
					Times(1)

				dbConn.
					EXPECT().
					Rollback(gomock.Eq(ctx)).
					Return(fmt.Errorf("rollback error")).
					Times(1)

				res, err := s.DeleteFile(ctx, p)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("rollback error")))
			})
		})

		When("failed remove disk file", func() {
			It("should return error", func() {
				fileRepo.
					EXPECT().
					GetConnection().
					Return(dbConn).
					Times(1)
				dbConn.
					EXPECT().
					Start(gomock.Eq(ctx)).
					Return(nil).
					Times(1)

				fileRepo.
					EXPECT().
					FindFile(gomock.Eq(ctx), gomock.Eq(findParam)).
					Return(findRes, nil).
					Times(1)

				fileRepo.
					EXPECT().
					DeleteFile(gomock.Eq(ctx), gomock.Eq(deleteParam)).
					Return(deleteRes, nil).
					Times(1)

				fileManager.
					EXPECT().
					RemoveFile(gomock.Eq(ctx), gomock.Eq(removeParam)).
					Return(nil, fmt.Errorf("r/w error")).
					Times(1)

				dbConn.
					EXPECT().
					Rollback(gomock.Eq(ctx)).
					Return(nil).
					Times(1)

				res, err := s.DeleteFile(ctx, p)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("r/w error")))
			})
		})

		When("failed commit transaction", func() {
			It("should return error", func() {
				fileRepo.
					EXPECT().
					GetConnection().
					Return(dbConn).
					Times(1)
				dbConn.
					EXPECT().
					Start(gomock.Eq(ctx)).
					Return(nil).
					Times(1)

				fileRepo.
					EXPECT().
					FindFile(gomock.Eq(ctx), gomock.Eq(findParam)).
					Return(findRes, nil).
					Times(1)

				fileRepo.
					EXPECT().
					DeleteFile(gomock.Eq(ctx), gomock.Eq(deleteParam)).
					Return(deleteRes, nil).
					Times(1)

				fileManager.
					EXPECT().
					RemoveFile(gomock.Eq(ctx), gomock.Eq(removeParam)).
					Return(removeRes, nil).
					Times(1)

				dbConn.
					EXPECT().
					Commit(gomock.Eq(ctx)).
					Return(fmt.Errorf("commit error")).
					Times(1)

				res, err := s.DeleteFile(ctx, p)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("commit error")))
			})
		})

		When("success delete file", func() {
			It("should return error", func() {
				fileRepo.
					EXPECT().
					GetConnection().
					Return(dbConn).
					Times(1)
				dbConn.
					EXPECT().
					Start(gomock.Eq(ctx)).
					Return(nil).
					Times(1)

				fileRepo.
					EXPECT().
					FindFile(gomock.Eq(ctx), gomock.Eq(findParam)).
					Return(findRes, nil).
					Times(1)

				fileRepo.
					EXPECT().
					DeleteFile(gomock.Eq(ctx), gomock.Eq(deleteParam)).
					Return(deleteRes, nil).
					Times(1)

				fileManager.
					EXPECT().
					RemoveFile(gomock.Eq(ctx), gomock.Eq(removeParam)).
					Return(removeRes, nil).
					Times(1)

				dbConn.
					EXPECT().
					Commit(gomock.Eq(ctx)).
					Return(nil).
					Times(1)

				res, err := s.DeleteFile(ctx, p)

				Expect(res).To(Equal(finalRes))
				Expect(err).To(BeNil())
			})
		})
	})
})
