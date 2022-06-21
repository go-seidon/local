package rest_app_test

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-seidon/local/internal/healthcheck"
	"github.com/go-seidon/local/internal/mock"
	rest_app "github.com/go-seidon/local/internal/rest-app"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("Handler Package", func() {

	Context("NotFoundHandler", func() {
		var (
			handler    http.HandlerFunc
			r          *http.Request
			w          *mock.MockResponseWriter
			log        *mock.MockLogger
			serializer *mock.MockSerializer
		)

		BeforeEach(func() {
			t := GinkgoT()
			r = &http.Request{}
			ctrl := gomock.NewController(t)
			w = mock.NewMockResponseWriter(ctrl)
			log = mock.NewMockLogger(ctrl)
			serializer = mock.NewMockSerializer(ctrl)
			handler = rest_app.NewNotFoundHandler(log, serializer)
		})

		When("success call the function", func() {
			It("should write response", func() {

				b := rest_app.ResponseBody{
					Code:    "ERROR",
					Message: "resource not found",
				}

				log.
					EXPECT().
					Debug("In function: NotFoundHandler").
					Times(1)
				log.
					EXPECT().
					Debug("Returning function: NotFoundHandler").
					Times(1)

				serializer.
					EXPECT().
					Encode(b).
					Return([]byte{}, nil).
					Times(1)

				w.
					EXPECT().
					Header().
					Return(http.Header{}).
					Times(1)

				w.
					EXPECT().
					WriteHeader(http.StatusNotFound).
					Times(1)

				w.
					EXPECT().
					Write([]byte{}).
					Times(1)

				handler.ServeHTTP(w, r)
			})
		})
	})

	Context("MethodNowAllowedHandler", func() {
		var (
			handler    http.HandlerFunc
			r          *http.Request
			w          *mock.MockResponseWriter
			log        *mock.MockLogger
			serializer *mock.MockSerializer
		)

		BeforeEach(func() {
			t := GinkgoT()
			r = &http.Request{}
			ctrl := gomock.NewController(t)
			w = mock.NewMockResponseWriter(ctrl)
			log = mock.NewMockLogger(ctrl)
			serializer = mock.NewMockSerializer(ctrl)
			handler = rest_app.NewMethodNotAllowedHandler(log, serializer)
		})

		When("success call the function", func() {
			It("should write response", func() {

				b := rest_app.ResponseBody{
					Code:    "ERROR",
					Message: "method is not allowed",
				}

				log.
					EXPECT().
					Debug("In function: MethodNotAllowedHandler").
					Times(1)
				log.
					EXPECT().
					Debug("Returning function: MethodNotAllowedHandler").
					Times(1)

				serializer.
					EXPECT().
					Encode(b).
					Return([]byte{}, nil).
					Times(1)

				w.
					EXPECT().
					Header().
					Return(http.Header{}).
					Times(1)

				w.
					EXPECT().
					WriteHeader(http.StatusMethodNotAllowed).
					Times(1)

				w.
					EXPECT().
					Write([]byte{}).
					Times(1)

				handler.ServeHTTP(w, r)
			})
		})
	})

	Context("RootHandler", func() {
		var (
			handler    http.HandlerFunc
			r          *http.Request
			w          *mock.MockResponseWriter
			log        *mock.MockLogger
			serializer *mock.MockSerializer
		)

		BeforeEach(func() {
			t := GinkgoT()
			r = &http.Request{}
			ctrl := gomock.NewController(t)
			w = mock.NewMockResponseWriter(ctrl)
			log = mock.NewMockLogger(ctrl)
			serializer = mock.NewMockSerializer(ctrl)
			handler = rest_app.NewRootHandler(log, serializer, "mock-name", "mock-version")
		})

		When("success call the function", func() {
			It("should write response", func() {

				b := rest_app.ResponseBody{
					Code:    "SUCCESS",
					Message: "success",
					Data: &rest_app.RootResult{
						AppName:    "mock-name",
						AppVersion: "mock-version",
					},
				}

				log.
					EXPECT().
					Debug("In function: RootHandler").
					Times(1)
				log.
					EXPECT().
					Debug("Returning function: RootHandler").
					Times(1)

				serializer.
					EXPECT().
					Encode(b).
					Return([]byte{}, nil).
					Times(1)

				w.
					EXPECT().
					Write([]byte{}).
					Times(1)

				handler.ServeHTTP(w, r)
			})
		})
	})

	Context("HealthCheckHandler", func() {
		var (
			handler       http.HandlerFunc
			r             *http.Request
			w             *mock.MockResponseWriter
			log           *mock.MockLogger
			serializer    *mock.MockSerializer
			healthService *mock.MockHealthService
		)

		BeforeEach(func() {
			t := GinkgoT()
			r = &http.Request{}
			ctrl := gomock.NewController(t)
			w = mock.NewMockResponseWriter(ctrl)
			log = mock.NewMockLogger(ctrl)
			serializer = mock.NewMockSerializer(ctrl)
			healthService = mock.NewMockHealthService(ctrl)
			handler = rest_app.NewHealthCheckHandler(log, serializer, healthService)
		})

		When("failed check service health", func() {
			It("should write response", func() {

				err := fmt.Errorf("failed check health")

				b := rest_app.ResponseBody{
					Code:    "ERROR",
					Message: err.Error(),
				}

				log.
					EXPECT().
					Debug("In function: HealthCheckHandler").
					Times(1)
				log.
					EXPECT().
					Debug("Returning function: HealthCheckHandler").
					Times(1)

				healthService.
					EXPECT().
					Check().
					Return(nil, err).
					Times(1)

				serializer.
					EXPECT().
					Encode(b).
					Return([]byte{}, nil).
					Times(1)

				w.
					EXPECT().
					WriteHeader(gomock.Eq(400)).
					Times(1)

				w.
					EXPECT().
					Write([]byte{}).
					Times(1)

				handler.ServeHTTP(w, r)
			})
		})

		When("success check service health", func() {
			It("should write response", func() {

				currentTimestamp := time.Now()
				res := &healthcheck.CheckResult{
					Status: "WARNING",
					Items: map[string]healthcheck.CheckResultItem{
						"app-disk": healthcheck.CheckResultItem{
							Name:      "app-disk",
							Status:    "FAILED",
							Error:     "Critical: disk usage too high 96.71 percent",
							CheckedAt: currentTimestamp,
							Metadata:  nil,
						},
						"internet-connection": healthcheck.CheckResultItem{
							Name:      "internet-connection",
							Status:    "OK",
							Error:     "",
							CheckedAt: currentTimestamp,
							Metadata:  nil,
						},
					},
				}
				jobs := map[string]rest_app.HealthCheckItem{
					"app-disk": rest_app.HealthCheckItem{
						Name:      "app-disk",
						Status:    "FAILED",
						Error:     "Critical: disk usage too high 96.71 percent",
						CheckedAt: currentTimestamp,
						Metadata:  nil,
					},
					"internet-connection": rest_app.HealthCheckItem{
						Name:      "internet-connection",
						Status:    "OK",
						Error:     "",
						CheckedAt: currentTimestamp,
						Metadata:  nil,
					},
				}

				b := rest_app.ResponseBody{
					Data: &rest_app.HealthCheckResponse{
						Status:  "WARNING",
						Details: jobs,
					},
					Code:    "SUCCESS",
					Message: "success check service health",
				}

				log.
					EXPECT().
					Debug("In function: HealthCheckHandler").
					Times(1)
				log.
					EXPECT().
					Debug("Returning function: HealthCheckHandler").
					Times(1)

				healthService.
					EXPECT().
					Check().
					Return(res, nil).
					Times(1)

				serializer.
					EXPECT().
					Encode(b).
					Return([]byte{}, nil).
					Times(1)

				w.
					EXPECT().
					Write([]byte{}).
					Times(1)

				handler.ServeHTTP(w, r)
			})
		})
	})
})
