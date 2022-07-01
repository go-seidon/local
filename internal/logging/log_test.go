package logging_test

import (
	"testing"

	"github.com/go-seidon/local/internal/logging"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLog(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Log Package")
}

var _ = Describe("Logging Package", func() {

	Context("WithAppContext function", Label("unit"), func() {
		When("parameter is specified", func() {
			It("should return result", func() {
				opt := logging.WithAppContext("mock-name", "mock-version")
				var res logging.LogOption
				opt(&res)

				Expect(res.AppName).To(Equal("mock-name"))
				Expect(res.AppVersion).To(Equal("mock-version"))
				Expect(res.DebuggingEnabled).To(BeFalse())
			})
		})
	})

	Context("EnableDebugging function", Label("unit"), func() {
		When("function is called", func() {
			It("should return result", func() {
				opt := logging.EnableDebugging()
				var res logging.LogOption
				opt(&res)

				Expect(res.DebuggingEnabled).To(BeTrue())
				Expect(res.AppName).To(Equal(""))
				Expect(res.AppVersion).To(Equal(""))
			})
		})
	})

})
