package logging_test

import (
	"github.com/go-seidon/local/internal/logging"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Logrus Package", func() {

	Context("NewLogrusLog function", Label("unit"), func() {
		When("option is not specified", func() {
			It("should return result", func() {
				res := logging.NewLogrusLog()

				Expect(res).ToNot(BeNil())
			})
		})

		When("option is specified", func() {
			It("should return result", func() {
				opt := logging.WithAppContext("mock-name", "mock-version")
				res := logging.NewLogrusLog(opt)

				Expect(res).ToNot(BeNil())
			})
		})
	})

	Context("Info function", Label("unit"), func() {
		var (
			logger logging.Logger
		)

		BeforeEach(func() {
			logger = logging.NewLogrusLog()
		})

		When("success send log", func() {
			It("should return nil", func() {
				logger.Info("mock-log")

			})
		})
	})

	Context("Debug function", Label("unit"), func() {
		var (
			logger logging.Logger
		)

		BeforeEach(func() {
			logger = logging.NewLogrusLog()
		})

		When("success send log", func() {
			It("should return nil", func() {
				logger.Debug("mock-log")

			})
		})
	})

	Context("Error function", Label("unit"), func() {
		var (
			logger logging.Logger
		)

		BeforeEach(func() {
			logger = logging.NewLogrusLog()
		})

		When("success send log", func() {
			It("should return nil", func() {
				logger.Error("mock-log")

			})
		})
	})

	Context("Warn function", Label("unit"), func() {
		var (
			logger logging.Logger
		)

		BeforeEach(func() {
			logger = logging.NewLogrusLog()
		})

		When("success send log", func() {
			It("should return nil", func() {
				logger.Warn("mock-log")

			})
		})
	})

})
