package rest_app

import (
	"net/http"
	"time"

	"github.com/go-seidon/local/internal/healthcheck"
	"github.com/go-seidon/local/internal/logging"
	"github.com/go-seidon/local/internal/serialization"
)

func NewNotFoundHandler(log logging.Logger, serializer serialization.Serializer) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Debug("In function: NotFoundHandler")
		defer log.Debug("Returning function: NotFoundHandler")

		w.Header().Set("Content-Type", "application/json")
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

		w.Header().Set("Content-Type", "application/json")
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

type HealthCheckItem struct {
	Name      string      `json:"name"`
	Status    string      `json:"status"`
	CheckedAt time.Time   `json:"checked_at"`
	Error     string      `json:"error"`
	Metadata  interface{} `json:"metadata"`
}

type HealthCheckResponse struct {
	Status  string                     `json:"status"`
	Details map[string]HealthCheckItem `json:"details"`
}

func NewHealthCheckHandler(log logging.Logger, serializer serialization.Serializer, healthService healthcheck.HealthService) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Debug("In function: HealthCheckHandler")
		defer log.Debug("Returning function: HealthCheckHandler")

		var res []byte
		var b ResponseBody

		r, err := healthService.Check()
		if err != nil {
			b = NewResponseBody(&NewResponseBodyParam{
				Code:    CODE_ERROR,
				Message: err.Error(),
			})

			w.WriteHeader(http.StatusBadRequest)
		} else {
			jobs := map[string]HealthCheckItem{}
			for jobName, item := range r.Items {
				jobs[jobName] = HealthCheckItem{
					Name:      item.Name,
					Status:    item.Status,
					Error:     item.Error,
					Metadata:  item.Metadata,
					CheckedAt: item.CheckedAt,
				}
			}

			b = NewResponseBody(&NewResponseBodyParam{
				Data: &HealthCheckResponse{
					Status:  r.Status,
					Details: jobs,
				},
				Message: "success check service health",
			})
		}

		res, _ = serializer.Encode(b)

		w.Write(res)
	}
}
