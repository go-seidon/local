package rest_app

import (
	"net/http"

	"github.com/go-seidon/local/internal/logging"
	"github.com/go-seidon/local/internal/serialization"
)

func NewNotFoundHandler(log logging.Logger, serializer serialization.Serializer) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Debug("In function: NotFoundHandler")
		defer log.Debug("Returning function: NotFoundHandler")

		w.Header().Set("Content-Type", "applicaton/json")
		w.WriteHeader(http.StatusNotFound)

		b := NewResponseBody(&NewResponseBodyParam{
			Code:    CODE_ERROR,
			Message: "resource not found",
		})
		res, _ := serializer.Encode(b)

		w.Write(res)
	}
}

func NewMethodNotAllowedHandler(log logging.Logger, serializer serialization.Serializer) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Debug("In function: MethodNotAllowedHandler")
		defer log.Debug("Returning function: MethodNotAllowedHandler")

		w.Header().Set("Content-Type", "applicaton/json")
		w.WriteHeader(http.StatusMethodNotAllowed)

		b := NewResponseBody(&NewResponseBodyParam{
			Code:    CODE_ERROR,
			Message: "method is not allowed",
		})
		res, _ := serializer.Encode(b)

		w.Write(res)
	}
}

type RootResult struct {
	AppName    string `json:"app_name"`
	AppVersion string `json:"app_version"`
}

func NewRootHandler(log logging.Logger, serializer serialization.Serializer, appName, appVersion string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Debug("In function: RootHandler")
		defer log.Debug("Returning function: RootHandler")

		b := NewResponseBody(&NewResponseBodyParam{
			Data: &RootResult{
				AppName:    appName,
				AppVersion: appVersion,
			},
		})
		res, _ := serializer.Encode(b)

		w.Write(res)
	}
}

type HealthCheckResponse struct {
	Status string `json:"status"`
}

func NewHealthCheckHandler(log logging.Logger, serializer serialization.Serializer) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Debug("In function: HealthCheckHandler")
		defer log.Debug("Returning function: HealthCheckHandler")

		b := NewResponseBody(&NewResponseBodyParam{
			Data: &HealthCheckResponse{
				Status: "ok",
			},
			Message: "success check service health",
		})
		res, _ := serializer.Encode(b)

		w.Write(res)
	}
}
