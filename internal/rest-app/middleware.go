package rest_app

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-seidon/local/internal/auth"
	"github.com/go-seidon/local/internal/serialization"
)

func DefaultHeaderMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		h.ServeHTTP(w, r)
	})
}

func NewBasicAuthMiddleware(a auth.BasicAuth, s serialization.Serializer) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authTokens := strings.Split(r.Header.Get("Authorization"), "Basic ")
			if len(authTokens) != 2 {
				Response(
					WithWriterSerializer(w, s),
					WithMessage("credential is not specified"),
					WithHttpCode(http.StatusUnauthorized),
					WithCode(CODE_UNAUTHORIZED),
				)
				return
			}

			res, err := a.CheckCredential(context.Background(), auth.CheckCredentialParam{
				AuthToken: authTokens[1],
			})
			if err != nil {
				Response(
					WithWriterSerializer(w, s),
					WithHttpCode(http.StatusUnauthorized),
					WithCode(CODE_UNAUTHORIZED),
					WithMessage("failed check credential"),
				)
				return
			}
			if !res.TokenValid {
				Response(
					WithWriterSerializer(w, s),
					WithMessage("credential is invalid"),
					WithHttpCode(http.StatusUnauthorized),
					WithCode(CODE_UNAUTHORIZED),
				)
				return
			}

			h.ServeHTTP(w, r)
		})
	}
}
