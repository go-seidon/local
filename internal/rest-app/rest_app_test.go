package rest_app_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/go-seidon/local/internal/logging"
	rest_app "github.com/go-seidon/local/internal/rest-app"
)

func TestRestApp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rest App Package")
}

var _ = Describe("Response Package", func() {

	Context("NewRestApp function", Label("unit"), func() {
		var (
			opt *rest_app.NewRestAppOption
		)

		BeforeEach(func() {
			log := logging.NewLogrusLog()
			opt = &rest_app.NewRestAppOption{
				Config: &rest_app.RestAppConfig{},
				Logger: log,
			}
		})

		When("options is not specified", func() {
			It("should return error", func() {
				opt = nil
				res, err := rest_app.NewRestApp(opt)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("invalid rest app option")))
			})
		})

		When("config is not specified", func() {
			It("should return error", func() {
				opt.Config = nil
				res, err := rest_app.NewRestApp(opt)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("invalid rest app config")))
			})
		})

		When("logger is not specified", func() {
			It("should return result", func() {
				opt.Logger = nil
				res, err := rest_app.NewRestApp(opt)

				Expect(res).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})

		When("parameter is specified", func() {
			It("should return result", func() {
				res, err := rest_app.NewRestApp(opt)

				Expect(res).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})

	Context("RestAppConfig", Label("unit"), func() {
		var (
			cfg *rest_app.RestAppConfig
		)

		BeforeEach(func() {
			cfg = &rest_app.RestAppConfig{
				AppName:    "mock-name",
				AppVersion: "mock-version",
				AppHost:    "host",
				AppPort:    3000,
			}
		})

		When("GetAppName function is called", func() {
			It("should return app name", func() {
				r := cfg.GetAppName()

				Expect(r).To(Equal("mock-name"))
			})
		})

		When("GetAppVersion function is called", func() {
			It("should return app name", func() {
				r := cfg.GetAppVersion()

				Expect(r).To(Equal("mock-version"))
			})
		})

		When("GetAddress function is called", func() {
			It("should return app name", func() {
				r := cfg.GetAddress()

				Expect(r).To(Equal("host:3000"))
			})
		})
	})
})
