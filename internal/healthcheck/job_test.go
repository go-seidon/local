package healthcheck_test

import (
	"fmt"
	"net/url"

	"github.com/go-seidon/local/internal/healthcheck"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Health Check Job", func() {

	Context("NewConnectionChecker function", Label("unit"), func() {
		var (
			p healthcheck.NewConnectionCheckerParam
		)

		BeforeEach(func() {
			p = healthcheck.NewConnectionCheckerParam{
				Url: "https://google.com",
			}
		})

		When("url is invalid", func() {
			It("should return error", func() {
				p.Url = "http:// "
				res, err := healthcheck.NewConnectionChecker(p)

				expectedErr := &url.Error{
					Op:  "parse",
					URL: "http:// ",
					Err: url.InvalidHostError(" "),
				}
				Expect(res).To(BeNil())
				Expect(err).To(Equal(expectedErr))
			})
		})

		When("parameter are valid", func() {
			It("should return result", func() {
				res, err := healthcheck.NewConnectionChecker(p)

				Expect(res).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})

	Context("NewDiskUsageChecker function", Label("unit"), func() {
		var (
			p healthcheck.NewDiskUsageCheckerParam
		)

		BeforeEach(func() {
			p = healthcheck.NewDiskUsageCheckerParam{
				Directory: "/usr/bin",
			}
		})

		When("directory is invalid", func() {
			It("should return error", func() {
				p.Directory = " "
				res, err := healthcheck.NewDiskUsageChecker(p)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("invalid directory")))
			})
		})

		When("parameter are valid", func() {
			It("should return result", func() {
				res, err := healthcheck.NewDiskUsageChecker(p)

				Expect(res).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})

	Context("NewHealthJobs function", Label("unit"), func() {

		When("parameter are valid", func() {
			It("should return result", func() {
				res, err := healthcheck.NewHealthJobs()

				Expect(res).ToNot(BeNil())
				Expect(len(res)).To(Equal(2))
				Expect(err).To(BeNil())
			})
		})
	})

})
