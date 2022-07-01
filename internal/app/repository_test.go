package app_test

import (
	"fmt"

	"github.com/go-seidon/local/internal/app"
	"github.com/go-seidon/local/internal/mock"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	_ "github.com/go-sql-driver/mysql"
)

var _ = Describe("Repository Package", func() {
	Context("NewRepository function", Label("unit"), func() {
		var (
			opt *mock.MockRepositoryOption
		)

		BeforeEach(func() {
			t := GinkgoT()
			ctrl := gomock.NewController(t)
			opt = mock.NewMockRepositoryOption(ctrl)
		})

		When("option is not specified", func() {
			It("should return error", func() {
				res, err := app.NewRepository(nil)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("invalid repository option")))
			})
		})

		When("provider is not supported", func() {
			It("should return error", func() {
				opt.EXPECT().Apply(&app.NewRepositoryOption{})
				res, err := app.NewRepository(opt)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("db provider is not supported")))
			})
		})

		When("success create mysql repository", func() {
			It("should return result", func() {
				opt := app.WithMySQLRepository("mock-username", "mock-password", "mock-db", "mock-host", 3306)
				res, err := app.NewRepository(opt)

				Expect(res).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
})
