package rest_app_test

import (
	"fmt"
	"net/http"

	"github.com/go-seidon/local/internal/mock"
	rest_app "github.com/go-seidon/local/internal/rest-app"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Response Package", func() {

	Context("WithWriterSerializer function", Label("unit"), func() {
		var (
			w *mock.MockResponseWriter
			s *mock.MockSerializer
		)

		BeforeEach(func() {
			t := GinkgoT()
			ctrl := gomock.NewController(t)
			w = mock.NewMockResponseWriter(ctrl)
			s = mock.NewMockSerializer(ctrl)
		})

		When("writer is not specified", func() {
			It("should not set writer", func() {
				p := rest_app.ResponseParam{}
				opt := rest_app.WithWriterSerializer(nil, s)

				opt(&p)

				Expect(p.Writer).To(BeNil())
				Expect(p.Serializer).To(Equal(s))
			})
		})

		When("serializer is not specified", func() {
			It("should not set serializer", func() {
				p := rest_app.ResponseParam{}
				opt := rest_app.WithWriterSerializer(w, nil)

				opt(&p)

				Expect(p.Writer).To(Equal(w))
				Expect(p.Serializer).To(BeNil())
			})
		})

		When("writer and serializer are specified", func() {
			It("should set writer and serializer", func() {
				p := rest_app.ResponseParam{}
				opt := rest_app.WithWriterSerializer(w, s)

				opt(&p)

				Expect(p.Writer).To(Equal(w))
				Expect(p.Serializer).To(Equal(s))
			})
		})
	})

	Context("WithHttpCode function", Label("unit"), func() {
		var (
			code int
		)

		BeforeEach(func() {
			code = http.StatusCreated
		})

		When("http code is specified", func() {
			It("should set http code", func() {
				p := rest_app.ResponseParam{}
				opt := rest_app.WithHttpCode(code)

				opt(&p)

				Expect(p.HttpCode).To(Equal(code))
			})
		})
	})

	Context("WithMessage function", Label("unit"), func() {
		var (
			message string
		)

		BeforeEach(func() {
			message = "success do something"
		})

		When("message is specified", func() {
			It("should set http", func() {
				p := rest_app.ResponseParam{}
				opt := rest_app.WithMessage(message)

				opt(&p)

				Expect(p.Message).To(Equal(message))
			})
		})
	})

	Context("Success function", Label("unit"), func() {
		var (
			data interface{}
		)

		BeforeEach(func() {
			data = struct{}{}
		})

		When("data is specified", func() {
			It("should set data", func() {
				p := rest_app.ResponseParam{}
				opt := rest_app.Success(data)

				opt(&p)

				Expect(p.Body.Code).To(Equal("SUCCESS"))
				Expect(p.Body.Message).To(Equal("success"))
				Expect(p.Body.Data).To(Equal(data))
			})
		})
	})

	Context("Error function", Label("unit"), func() {
		var (
			message string
		)

		BeforeEach(func() {
			message = "failed do something"
		})

		When("message is specified", func() {
			It("should set message", func() {
				p := rest_app.ResponseParam{}
				opt := rest_app.Error(message)

				opt(&p)

				Expect(p.Body.Code).To(Equal("ERROR"))
				Expect(p.Body.Message).To(Equal(message))
				Expect(p.Body.Data).To(BeNil())
			})
		})

		When("message is not specified", func() {
			It("should use default message", func() {
				p := rest_app.ResponseParam{}
				opt := rest_app.Error("")

				opt(&p)

				Expect(p.Body.Code).To(Equal("ERROR"))
				Expect(p.Body.Message).To(Equal("error"))
				Expect(p.Body.Data).To(BeNil())
			})
		})
	})

	Context("NotFound function", Label("unit"), func() {
		var (
			message string
		)

		BeforeEach(func() {
			message = "something is not found"
		})

		When("message is specified", func() {
			It("should set message", func() {
				p := rest_app.ResponseParam{}
				opt := rest_app.NotFound(message)

				opt(&p)

				Expect(p.Body.Code).To(Equal("NOT_FOUND"))
				Expect(p.Body.Message).To(Equal(message))
				Expect(p.Body.Data).To(BeNil())
			})
		})

		When("message is not specified", func() {
			It("should use default message", func() {
				p := rest_app.ResponseParam{}
				opt := rest_app.NotFound("")

				opt(&p)

				Expect(p.Body.Code).To(Equal("NOT_FOUND"))
				Expect(p.Body.Message).To(Equal("not found"))
				Expect(p.Body.Data).To(BeNil())
			})
		})
	})

	Context("Response function", Label("unit"), func() {
		var (
			w *mock.MockResponseWriter
			s *mock.MockSerializer
		)

		BeforeEach(func() {
			t := GinkgoT()
			ctrl := gomock.NewController(t)
			w = mock.NewMockResponseWriter(ctrl)
			s = mock.NewMockSerializer(ctrl)
		})

		When("writer is not specified", func() {
			It("should return error", func() {
				err := rest_app.Response(
					rest_app.WithWriterSerializer(nil, s),
				)

				Expect(err).To(Equal(fmt.Errorf("writer should be specified")))
			})
		})

		When("serializer is not specified", func() {
			It("should return error", func() {
				err := rest_app.Response(
					rest_app.WithWriterSerializer(w, nil),
				)

				Expect(err).To(Equal(fmt.Errorf("serializer should be specified")))
			})
		})

		When("http code is specified", func() {
			It("should return nil", func() {
				s.
					EXPECT().
					Encode(gomock.Any()).
					Return([]byte("mock"), nil).
					Times(1)

				w.
					EXPECT().
					WriteHeader(gomock.Eq(201)).
					Times(1)

				w.
					EXPECT().
					Write(gomock.Eq([]byte("mock"))).
					Times(1)

				err := rest_app.Response(
					rest_app.WithWriterSerializer(w, s),
					rest_app.WithHttpCode(201),
				)

				Expect(err).To(BeNil())
			})
		})

		When("message is specified", func() {
			It("should return nil", func() {
				b := rest_app.ResponseBody{
					Code:    "SUCCESS",
					Message: "success do something",
				}
				s.
					EXPECT().
					Encode(gomock.Eq(b)).
					Return([]byte("mock"), nil).
					Times(1)

				w.
					EXPECT().
					WriteHeader(gomock.Eq(200)).
					Times(1)

				w.
					EXPECT().
					Write(gomock.Eq([]byte("mock"))).
					Times(1)

				err := rest_app.Response(
					rest_app.WithWriterSerializer(w, s),
					rest_app.WithMessage("success do something"),
				)

				Expect(err).To(BeNil())
			})
		})

		When("failed encode data", func() {
			It("should return error", func() {
				s.
					EXPECT().
					Encode(gomock.Any()).
					Return(nil, fmt.Errorf("failed encode")).
					Times(1)

				err := rest_app.Response(
					rest_app.WithWriterSerializer(w, s),
				)

				Expect(err).To(Equal(fmt.Errorf("failed encode")))
			})
		})

		When("success encode data", func() {
			It("should return nil", func() {
				s.
					EXPECT().
					Encode(gomock.Any()).
					Return([]byte("mock"), nil).
					Times(1)

				w.
					EXPECT().
					WriteHeader(gomock.Eq(200)).
					Times(1)

				w.
					EXPECT().
					Write(gomock.Eq([]byte("mock"))).
					Times(1)

				err := rest_app.Response(
					rest_app.WithWriterSerializer(w, s),
				)

				Expect(err).To(BeNil())
			})
		})

	})

})
