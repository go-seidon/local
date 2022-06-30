package rest_app_test

import (
	"net/http"

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

})
