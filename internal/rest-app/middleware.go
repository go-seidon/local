package rest_app

import (
	"net/http"
)

func DefaultHeaderMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "applicaton/json")
		w.WriteHeader(http.StatusOK)

		h.ServeHTTP(w, r)
	})
}
