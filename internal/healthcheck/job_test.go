package healthcheck_test

import (
	"fmt"
	"net/url"
	"time"

	"github.com/go-seidon/local/internal/healthcheck"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Health Check Job", func() {

	Context("NewHttpPingJob function", Label("unit"), func() {
		var (
			p healthcheck.NewHttpPingJobParam
		)

		BeforeEach(func() {
			p = healthcheck.NewHttpPingJobParam{
				Name:     "internet-checker",
				Interval: 30 * time.Second,
				Url:      "https://google.com",
			}
		})

		When("name is invalid", func() {
			It("should return error", func() {
				p.Name = " "
				res, err := healthcheck.NewHttpPingJob(p)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("invalid name")))
			})
		})

		When("url is invalid", func() {
			It("should return error", func() {
				p.Url = "http:// "
				res, err := healthcheck.NewHttpPingJob(p)

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
				res, err := healthcheck.NewHttpPingJob(p)

				Expect(res).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})

	Context("NewDiskUsageJob function", Label("unit"), func() {
		var (
			p healthcheck.NewDiskUsageJobParam
		)

		BeforeEach(func() {
			p = healthcheck.NewDiskUsageJobParam{
				Name:      "app-disk",
				Interval:  60 * time.Second,
				Directory: "/usr/bin",
			}
		})

		When("name is invalid", func() {
			It("should return error", func() {
				p.Name = " "
				res, err := healthcheck.NewDiskUsageJob(p)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("invalid name")))
			})
		})

		When("directory is invalid", func() {
			It("should return error", func() {
				p.Directory = " "
				res, err := healthcheck.NewDiskUsageJob(p)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("invalid directory")))
			})
		})

		When("parameter are valid", func() {
			It("should return result", func() {
				res, err := healthcheck.NewDiskUsageJob(p)

				Expect(res).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})

})
