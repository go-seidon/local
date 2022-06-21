package rest_app_test

import (
	rest_app "github.com/go-seidon/local/internal/rest-app"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Response Package", func() {

	Context("NewResponseBody function", func() {
		var (
			p *rest_app.NewResponseBodyParam
		)

		BeforeEach(func() {
			p = &rest_app.NewResponseBodyParam{
				Code:    "CUSTOM-CODE",
				Message: "custom-message",
				Data:    struct{}{},
			}
		})

		When("parameter is not specified", func() {
			It("should return default result", func() {
				p = nil
				b := rest_app.NewResponseBody(p)

				Expect(b.Code).To(Equal("SUCCESS"))
				Expect(b.Message).To(Equal("success"))
				Expect(b.Data).To(BeNil())
			})
		})

		When("code is specified", func() {
			It("should return result", func() {
				p.Message = ""
				p.Data = nil
				b := rest_app.NewResponseBody(p)

				Expect(b.Code).To(Equal("CUSTOM-CODE"))
				Expect(b.Message).To(Equal("success"))
				Expect(b.Data).To(BeNil())
			})
		})

		When("message is specified", func() {
			It("should return result", func() {
				p.Code = ""
				p.Data = nil
				b := rest_app.NewResponseBody(p)

				Expect(b.Code).To(Equal("SUCCESS"))
				Expect(b.Message).To(Equal("custom-message"))
				Expect(b.Data).To(BeNil())
			})
		})

		When("data is specified", func() {
			It("should return result", func() {
				p.Code = ""
				p.Message = ""
				b := rest_app.NewResponseBody(p)

				Expect(b.Code).To(Equal("SUCCESS"))
				Expect(b.Message).To(Equal("success"))
				Expect(b.Data).To(Equal(struct{}{}))
			})
		})
	})

})
