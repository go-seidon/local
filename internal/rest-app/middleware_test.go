package rest_app_test

import (
	"fmt"
	"net/http"

	"github.com/go-seidon/local/internal/auth"
	"github.com/go-seidon/local/internal/mock"
	rest_app "github.com/go-seidon/local/internal/rest-app"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("Middleware Package", func() {

	Context("DefaultHeaderMiddleware", Label("unit"), func() {
		var (
			r           *http.Request
			w           *mock.MockResponseWriter
			middleware  http.Handler
			httpHandler *mock.MockHandler
		)

		BeforeEach(func() {
			t := GinkgoT()
			ctrl := gomock.NewController(t)

			r = &http.Request{}
			w = mock.NewMockResponseWriter(ctrl)
			httpHandler = mock.NewMockHandler(ctrl)
			middleware = rest_app.DefaultHeaderMiddleware(httpHandler)
		})

		When("middleware is called", func() {
			It("should call serve http", func() {
				httpHandler.EXPECT().
					ServeHTTP(gomock.Eq(w), gomock.Eq(r)).
					Times(1)

				w.EXPECT().
					Header().
					Return(http.Header{}).
					Times(1)

				middleware.ServeHTTP(w, r)
			})
		})
	})

	Context("NewBasicAuthMiddleware", Label("unit"), func() {
		var (
			a       *mock.MockBasicAuth
			s       *mock.MockSerializer
			handler *mock.MockHandler
			m       http.Handler

			rw  *mock.MockResponseWriter
			req *http.Request
		)

		BeforeEach(func() {
			t := GinkgoT()
			ctrl := gomock.NewController(t)
			a = mock.NewMockBasicAuth(ctrl)
			s = mock.NewMockSerializer(ctrl)
			handler = mock.NewMockHandler(ctrl)
			fn := rest_app.NewBasicAuthMiddleware(a, s)
			m = fn(handler)

			rw = mock.NewMockResponseWriter(ctrl)
			req = &http.Request{
				Header: http.Header{},
			}
			req.Header.Set("Authorization", "Basic basic-token")
		})

		When("basic auth is not specified", func() {
			It("should return error", func() {
				req.Header.Del("Authorization")

				b := rest_app.ResponseBody{
					Code:    "UNAUTHORIZED",
					Message: "credential is not specified",
				}
				s.
					EXPECT().
					Marshal(gomock.Eq(b)).
					Return([]byte{}, nil).
					Times(1)
				rw.
					EXPECT().
					WriteHeader(401).
					Times(1)
				rw.
					EXPECT().
					Write(gomock.Eq([]byte{})).
					Times(1)

				m.ServeHTTP(rw, req)
			})
		})

		When("failed check credential", func() {
			It("should return error", func() {
				checkParam := auth.CheckCredentialParam{
					AuthToken: "basic-token",
				}
				a.
					EXPECT().
					CheckCredential(gomock.Any(), gomock.Eq(checkParam)).
					Return(nil, fmt.Errorf("db error")).
					Times(1)

				b := rest_app.ResponseBody{
					Code:    "UNAUTHORIZED",
					Message: "failed check credential",
				}
				s.
					EXPECT().
					Marshal(gomock.Eq(b)).
					Return([]byte{}, nil).
					Times(1)
				rw.
					EXPECT().
					WriteHeader(401).
					Times(1)
				rw.
					EXPECT().
					Write(gomock.Eq([]byte{})).
					Times(1)

				m.ServeHTTP(rw, req)
			})
		})

		When("failed token is invalid", func() {
			It("should return error", func() {
				checkParam := auth.CheckCredentialParam{
					AuthToken: "basic-token",
				}
				checkRes := &auth.CheckCredentialResult{
					TokenValid: false,
				}
				a.
					EXPECT().
					CheckCredential(gomock.Any(), gomock.Eq(checkParam)).
					Return(checkRes, nil).
					Times(1)

				b := rest_app.ResponseBody{
					Code:    "UNAUTHORIZED",
					Message: "credential is invalid",
				}
				s.
					EXPECT().
					Marshal(gomock.Eq(b)).
					Return([]byte{}, nil).
					Times(1)
				rw.
					EXPECT().
					WriteHeader(401).
					Times(1)
				rw.
					EXPECT().
					Write(gomock.Eq([]byte{})).
					Times(1)

				m.ServeHTTP(rw, req)
			})
		})

		When("token is valid", func() {
			It("should return result", func() {
				checkParam := auth.CheckCredentialParam{
					AuthToken: "basic-token",
				}
				checkRes := &auth.CheckCredentialResult{
					TokenValid: true,
				}
				a.
					EXPECT().
					CheckCredential(gomock.Any(), gomock.Eq(checkParam)).
					Return(checkRes, nil).
					Times(1)

				handler.
					EXPECT().
					ServeHTTP(gomock.Eq(rw), gomock.Eq(req)).
					Times(1)

				m.ServeHTTP(rw, req)
			})
		})
	})

})
