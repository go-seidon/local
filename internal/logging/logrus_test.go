package logging_test

import (
	"fmt"
	"testing"

	"github.com/go-seidon/local/internal/logging"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLog(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Log Package")
}

var _ = Describe("Logger Package", func() {

	Context("NewLogrusLog function", func() {
		var (
			opt *logging.NewLogrusLogOption
		)

		When("option is not specified", func() {
			It("should return error", func() {
				res, err := logging.NewLogrusLog(opt)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("logrus option is invalid")))
			})
		})

		When("all parameter is specified", func() {
			It("should return result", func() {
				opt = &logging.NewLogrusLogOption{
					AppName:    "mock-name",
					AppVersion: "mock-version",
				}
				res, err := logging.NewLogrusLog(opt)

				Expect(res).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})

	Context("Info function", func() {
		var (
			logger logging.Logger
		)

		BeforeEach(func() {
			logger, _ = logging.NewLogrusLog(&logging.NewLogrusLogOption{
				AppName:    "mock-name",
				AppVersion: "mock-version",
			})
		})

		When("success send log", func() {
			It("should return nil", func() {
				err := logger.Info("mock-log")

				Expect(err).To(BeNil())
			})
		})
	})

	Context("Debug function", func() {
		var (
			logger logging.Logger
		)

		BeforeEach(func() {
			logger, _ = logging.NewLogrusLog(&logging.NewLogrusLogOption{
				AppName:    "mock-name",
				AppVersion: "mock-version",
			})
		})

		When("success send log", func() {
			It("should return nil", func() {
				err := logger.Debug("mock-log")

				Expect(err).To(BeNil())
			})
		})
	})

})
