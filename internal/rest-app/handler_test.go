package rest_app_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-seidon/local/internal/deleting"
	"github.com/go-seidon/local/internal/healthcheck"
	"github.com/go-seidon/local/internal/mock"
	rest_app "github.com/go-seidon/local/internal/rest-app"
	"github.com/go-seidon/local/internal/retrieving"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("Handler Package", func() {

	Context("NotFoundHandler", Label("unit"), func() {
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
					Code:    "NOT_FOUND",
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

	Context("MethodNowAllowedHandler", Label("unit"), func() {
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

	Context("RootHandler", Label("unit"), func() {
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
					WriteHeader(gomock.Eq(200))

				w.
					EXPECT().
					Write([]byte{}).
					Times(1)

				handler.ServeHTTP(w, r)
			})
		})
	})

	Context("HealthCheckHandler", Label("unit"), func() {
		var (
			handler       http.HandlerFunc
			r             *http.Request
			w             *mock.MockResponseWriter
			log           *mock.MockLogger
			serializer    *mock.MockSerializer
			healthService *mock.MockHealthCheck
		)

		BeforeEach(func() {
			t := GinkgoT()
			r = &http.Request{}
			ctrl := gomock.NewController(t)
			w = mock.NewMockResponseWriter(ctrl)
			log = mock.NewMockLogger(ctrl)
			serializer = mock.NewMockSerializer(ctrl)
			healthService = mock.NewMockHealthCheck(ctrl)
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
						"app-disk": {
							Name:      "app-disk",
							Status:    "FAILED",
							Error:     "Critical: disk usage too high 96.71 percent",
							CheckedAt: currentTimestamp,
							Metadata:  nil,
						},
						"internet-connection": {
							Name:      "internet-connection",
							Status:    "OK",
							Error:     "",
							CheckedAt: currentTimestamp,
							Metadata:  nil,
						},
					},
				}
				jobs := map[string]rest_app.HealthCheckItem{
					"app-disk": {
						Name:      "app-disk",
						Status:    "FAILED",
						Error:     "Critical: disk usage too high 96.71 percent",
						CheckedAt: currentTimestamp,
						Metadata:  nil,
					},
					"internet-connection": {
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
					WriteHeader(gomock.Eq(200))

				w.
					EXPECT().
					Write([]byte{}).
					Times(1)

				handler.ServeHTTP(w, r)
			})
		})
	})

	Context("NewDeleteFileHandler", Label("unit"), func() {
		var (
			handler       http.HandlerFunc
			r             *http.Request
			w             *mock.MockResponseWriter
			log           *mock.MockLogger
			serializer    *mock.MockSerializer
			deleteService *mock.MockDeleter
			p             deleting.DeleteFileParam
		)

		BeforeEach(func() {
			t := GinkgoT()
			r = mux.SetURLVars(&http.Request{}, map[string]string{
				"unique_id": "mock-file-id",
			})
			ctrl := gomock.NewController(t)
			w = mock.NewMockResponseWriter(ctrl)
			log = mock.NewMockLogger(ctrl)
			serializer = mock.NewMockSerializer(ctrl)
			deleteService = mock.NewMockDeleter(ctrl)
			handler = rest_app.NewDeleteFileHandler(log, serializer, deleteService)
			p = deleting.DeleteFileParam{
				FileId: "mock-file-id",
			}
		})

		When("failed delete file", func() {
			It("should write response", func() {

				err := fmt.Errorf("failed delete file")

				b := rest_app.ResponseBody{
					Code:    "ERROR",
					Message: err.Error(),
				}

				log.
					EXPECT().
					Debug("In function: DeleteFileHandler").
					Times(1)
				log.
					EXPECT().
					Debug("Returning function: DeleteFileHandler").
					Times(1)

				deleteService.
					EXPECT().
					DeleteFile(gomock.Any(), gomock.Eq(p)).
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

		When("file is not found", func() {
			It("should write response", func() {

				err := deleting.ErrorResourceNotFound

				b := rest_app.ResponseBody{
					Code:    "NOT_FOUND",
					Message: err.Error(),
				}

				log.
					EXPECT().
					Debug("In function: DeleteFileHandler").
					Times(1)
				log.
					EXPECT().
					Debug("Returning function: DeleteFileHandler").
					Times(1)

				deleteService.
					EXPECT().
					DeleteFile(gomock.Any(), gomock.Eq(p)).
					Return(nil, err).
					Times(1)

				serializer.
					EXPECT().
					Encode(b).
					Return([]byte{}, nil).
					Times(1)

				w.
					EXPECT().
					WriteHeader(gomock.Eq(404)).
					Times(1)

				w.
					EXPECT().
					Write([]byte{}).
					Times(1)

				handler.ServeHTTP(w, r)
			})
		})

		When("success delete file", func() {
			It("should write response", func() {
				res := &deleting.DeleteFileResult{
					DeletedAt: time.Now(),
				}
				b := rest_app.ResponseBody{
					Code:    "SUCCESS",
					Message: "success delete file",
					Data: &rest_app.DeleteFileResponse{
						DeletedAt: res.DeletedAt.UnixMilli(),
					},
				}

				log.
					EXPECT().
					Debug("In function: DeleteFileHandler").
					Times(1)
				log.
					EXPECT().
					Debug("Returning function: DeleteFileHandler").
					Times(1)

				deleteService.
					EXPECT().
					DeleteFile(gomock.Any(), gomock.Eq(p)).
					Return(res, nil).
					Times(1)

				serializer.
					EXPECT().
					Encode(b).
					Return([]byte{}, nil).
					Times(1)

				w.
					EXPECT().
					WriteHeader(gomock.Eq(200)).
					Times(1)

				w.
					EXPECT().
					Write([]byte{}).
					Times(1)

				handler.ServeHTTP(w, r)
			})
		})
	})

	Context("NewRetrieveFileHandler", Label("unit"), func() {
		var (
			ctx             context.Context
			handler         http.HandlerFunc
			r               *http.Request
			w               *mock.MockResponseWriter
			log             *mock.MockLogger
			serializer      *mock.MockSerializer
			retrieveService *mock.MockRetriever
			fileData        *mock.MockReadCloser
			p               retrieving.RetrieveFileParam
		)

		BeforeEach(func() {
			t := GinkgoT()
			ctx = context.Background()
			r = mux.SetURLVars(&http.Request{}, map[string]string{
				"unique_id": "mock-file-id",
			})
			ctrl := gomock.NewController(t)
			w = mock.NewMockResponseWriter(ctrl)
			log = mock.NewMockLogger(ctrl)
			serializer = mock.NewMockSerializer(ctrl)
			retrieveService = mock.NewMockRetriever(ctrl)
			fileData = mock.NewMockReadCloser(ctrl)
			handler = rest_app.NewRetrieveFileHandler(log, serializer, retrieveService)
			p = retrieving.RetrieveFileParam{
				FileId: "mock-file-id",
			}
		})

		When("failed retrieve file", func() {
			It("should write response", func() {

				err := fmt.Errorf("failed retrieve file")

				b := rest_app.ResponseBody{
					Code:    "ERROR",
					Message: err.Error(),
				}

				log.
					EXPECT().
					Debug("In function: RetrieveFileHandler").
					Times(1)
				log.
					EXPECT().
					Debug("Returning function: RetrieveFileHandler").
					Times(1)

				retrieveService.
					EXPECT().
					RetrieveFile(gomock.Eq(ctx), gomock.Eq(p)).
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

		When("file is not found", func() {
			It("should write response", func() {

				err := retrieving.ErrorResourceNotFound

				b := rest_app.ResponseBody{
					Code:    "NOT_FOUND",
					Message: err.Error(),
				}

				log.
					EXPECT().
					Debug("In function: RetrieveFileHandler").
					Times(1)
				log.
					EXPECT().
					Debug("Returning function: RetrieveFileHandler").
					Times(1)

				retrieveService.
					EXPECT().
					RetrieveFile(gomock.Eq(ctx), gomock.Eq(p)).
					Return(nil, err).
					Times(1)

				serializer.
					EXPECT().
					Encode(b).
					Return([]byte{}, nil).
					Times(1)

				w.
					EXPECT().
					WriteHeader(gomock.Eq(404)).
					Times(1)

				w.
					EXPECT().
					Write([]byte{}).
					Times(1)

				handler.ServeHTTP(w, r)
			})
		})

		When("failed read file", func() {
			It("should write response", func() {

				fileData.
					EXPECT().
					Close().
					Times(1)

				fileData.
					EXPECT().
					Read(gomock.Any()).
					Return(0, fmt.Errorf("read error")).
					Times(1)

				res := &retrieving.RetrieveFileResult{
					Data: fileData,
				}

				b := rest_app.ResponseBody{
					Code:    "ERROR",
					Message: "read error",
				}

				log.
					EXPECT().
					Debug("In function: RetrieveFileHandler").
					Times(1)
				log.
					EXPECT().
					Debug("Returning function: RetrieveFileHandler").
					Times(1)

				retrieveService.
					EXPECT().
					RetrieveFile(gomock.Eq(ctx), gomock.Eq(p)).
					Return(res, nil).
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

		When("mimetype is empty", func() {
			It("should write response", func() {

				fileData.
					EXPECT().
					Close().
					Times(1)

				fileData.
					EXPECT().
					Read(gomock.Any()).
					Return(0, io.EOF).
					Times(1)

				res := &retrieving.RetrieveFileResult{
					Data:      fileData,
					UniqueId:  "mock-unique-id",
					Name:      "mock-name",
					Path:      "mock-path",
					MimeType:  "",
					Extension: "mock-extension",
					DeletedAt: nil,
				}

				log.
					EXPECT().
					Debug("In function: RetrieveFileHandler").
					Times(1)
				log.
					EXPECT().
					Debug("Returning function: RetrieveFileHandler").
					Times(1)

				retrieveService.
					EXPECT().
					RetrieveFile(gomock.Eq(ctx), gomock.Eq(p)).
					Return(res, nil).
					Times(1)

				w.EXPECT().
					Header().
					Return(http.Header{}).
					Times(1)

				w.
					EXPECT().
					Write([]byte{}).
					Times(1)

				handler.ServeHTTP(w, r)
			})
		})

		When("mimetype is not empty", func() {
			It("should write response", func() {

				fileData.
					EXPECT().
					Close().
					Times(1)

				fileData.
					EXPECT().
					Read(gomock.Any()).
					Return(0, io.EOF).
					Times(1)

				res := &retrieving.RetrieveFileResult{
					Data:      fileData,
					UniqueId:  "mock-unique-id",
					Name:      "mock-name",
					Path:      "mock-path",
					MimeType:  "text/plain",
					Extension: "mock-extension",
					DeletedAt: nil,
				}

				log.
					EXPECT().
					Debug("In function: RetrieveFileHandler").
					Times(1)
				log.
					EXPECT().
					Debug("Returning function: RetrieveFileHandler").
					Times(1)

				retrieveService.
					EXPECT().
					RetrieveFile(gomock.Eq(ctx), gomock.Eq(p)).
					Return(res, nil).
					Times(1)

				w.EXPECT().
					Header().
					Return(http.Header{}).
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
