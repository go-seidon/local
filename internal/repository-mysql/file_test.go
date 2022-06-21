package repository_mysql_test

import (
	"database/sql"
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	repository_mysql "github.com/go-seidon/local/internal/repository-mysql"
)

func TestRepositoryMySQL(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repository MySQL Package")
}

var _ = Describe("File Repository", func() {
	Context("NewFileRepository function", func() {
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
})
