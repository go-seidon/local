package rest_app

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-seidon/local/internal/deleting"
	"github.com/go-seidon/local/internal/healthcheck"
	"github.com/go-seidon/local/internal/logging"
	"github.com/go-seidon/local/internal/retrieving"
	"github.com/go-seidon/local/internal/serialization"
	"github.com/go-seidon/local/internal/uploading"
	"github.com/gorilla/mux"
)

func NewNotFoundHandler(log logging.Logger, serializer serialization.Serializer) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Debug("In function: NotFoundHandler")
		defer log.Debug("Returning function: NotFoundHandler")

		b := NewResponseBody(&NewResponseBodyParam{
			Code:    CODE_NOT_FOUND,
			Message: "resource not found",
		})
		res, _ := serializer.Encode(b)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(res)
	}
}

func NewMethodNotAllowedHandler(log logging.Logger, serializer serialization.Serializer) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Debug("In function: MethodNotAllowedHandler")
		defer log.Debug("Returning function: MethodNotAllowedHandler")

		b := NewResponseBody(&NewResponseBodyParam{
			Code:    CODE_ERROR,
			Message: "method is not allowed",
		})
		res, _ := serializer.Encode(b)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
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

		w.WriteHeader(http.StatusOK)
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

func NewHealthCheckHandler(log logging.Logger, serializer serialization.Serializer, healthService healthcheck.HealthCheck) http.HandlerFunc {
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

			w.WriteHeader(http.StatusOK)
		}

		res, _ = serializer.Encode(b)

		w.Write(res)
	}
}

type DeleteFileResponse struct {
	DeletedAt int64 `json:"deleted_at"`
}

func NewDeleteFileHandler(log logging.Logger, serializer serialization.Serializer, deleter deleting.Deleter) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Debug("In function: DeleteFileHandler")
		defer log.Debug("Returning function: DeleteFileHandler")

		var res []byte
		var b ResponseBody

		vars := mux.Vars(req)

		ctx := context.Background()
		r, err := deleter.DeleteFile(ctx, deleting.DeleteFileParam{
			FileId: vars["unique_id"],
		})
		if err == nil {
			b = NewResponseBody(&NewResponseBodyParam{
				Data: &DeleteFileResponse{
					DeletedAt: r.DeletedAt.UnixMilli(),
				},
				Message: "success delete file",
			})

			res, _ = serializer.Encode(b)
			w.WriteHeader(http.StatusOK)
			w.Write(res)

			return
		}

		if errors.Is(err, deleting.ErrorResourceNotFound) {
			b = NewResponseBody(&NewResponseBodyParam{
				Code:    CODE_NOT_FOUND,
				Message: err.Error(),
			})

			w.WriteHeader(http.StatusNotFound)
		} else {
			b = NewResponseBody(&NewResponseBodyParam{
				Code:    CODE_ERROR,
				Message: err.Error(),
			})

			w.WriteHeader(http.StatusBadRequest)
		}

		res, _ = serializer.Encode(b)
		w.Write(res)
	}
}

func NewRetrieveFileHandler(log logging.Logger, serializer serialization.Serializer, retriever retrieving.Retriever) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Debug("In function: RetrieveFileHandler")
		defer log.Debug("Returning function: RetrieveFileHandler")

		var res []byte
		var b ResponseBody

		vars := mux.Vars(req)

		ctx := context.Background()
		r, err := retriever.RetrieveFile(ctx, retrieving.RetrieveFileParam{
			FileId: vars["unique_id"],
		})
		if err == nil {

			defer r.Data.Close()
			data, err := io.ReadAll(r.Data)
			if err != nil {
				b = NewResponseBody(&NewResponseBodyParam{
					Code:    CODE_ERROR,
					Message: err.Error(),
				})
				res, _ = serializer.Encode(b)

				w.WriteHeader(http.StatusBadRequest)
				w.Write(res)
				return
			}

			if r.MimeType != "" {
				w.Header().Set("Content-Type", r.MimeType)
			} else {
				w.Header().Del("Content-Type")
			}

			w.Write(data)
			return
		}

		if errors.Is(err, retrieving.ErrorResourceNotFound) {
			b = NewResponseBody(&NewResponseBodyParam{
				Code:    CODE_NOT_FOUND,
				Message: err.Error(),
			})

			w.WriteHeader(http.StatusNotFound)
		} else {
			b = NewResponseBody(&NewResponseBodyParam{
				Code:    CODE_ERROR,
				Message: err.Error(),
			})

			w.WriteHeader(http.StatusBadRequest)
		}

		res, _ = serializer.Encode(b)
		w.Write(res)
	}
}

func NewUploadFileHandler(log logging.Logger, serializer serialization.Serializer, uploader uploading.Uploader, locator uploading.UploadLocation, config *RestAppConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Debug("In function: UploadFileHandler")
		defer log.Debug("Returning function: UploadFileHandler")

		// set form max size + add 1KB (non file size estimation if any)
		req.Body = http.MaxBytesReader(w, req.Body, config.UploadFormSize+1024)

		file, fileHeader, err := req.FormFile("file")
		if err != nil {
			b := NewResponseBody(&NewResponseBodyParam{
				Code:    CODE_ERROR,
				Message: err.Error(),
			})

			res, _ := serializer.Encode(b)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(res)

			return
		}
		defer file.Close()

		fileInfo, err := ParseMultipartFile(file, fileHeader)
		if err != nil {
			b := NewResponseBody(&NewResponseBodyParam{
				Code:    CODE_ERROR,
				Message: err.Error(),
			})

			res, _ := serializer.Encode(b)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(res)

			return
		}

		uploadDir := fmt.Sprintf("%s/%s", config.UploadDir, locator.GetLocation())

		ctx := context.Background()
		uploadRes, err := uploader.UploadFile(ctx,
			uploading.WithReader(file),
			uploading.WithDirectory(uploadDir),
			uploading.WithFileInfo(
				fileInfo.Name,
				fileInfo.Mimetype,
				fileInfo.Extension,
				fileInfo.Size,
			),
		)
		if err != nil {
			b := NewResponseBody(&NewResponseBodyParam{
				Code:    CODE_ERROR,
				Message: err.Error(),
			})

			res, _ := serializer.Encode(b)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(res)

			return
		}

		b := NewResponseBody(&NewResponseBodyParam{
			Data: struct {
				UniqueId   string `json:"id"`
				Name       string `json:"name"`
				Mimetype   string `json:"mimetype"`
				Extension  string `json:"extension"`
				Size       int64  `json:"size"`
				UploadedAt int64  `json:"uploaded_at"`
			}{
				UniqueId:   uploadRes.UniqueId,
				Name:       uploadRes.Name,
				Mimetype:   uploadRes.Mimetype,
				Extension:  uploadRes.Extension,
				Size:       uploadRes.Size,
				UploadedAt: uploadRes.UploadedAt.UnixMilli(),
			},
			Message: "success upload file",
		})

		res, _ := serializer.Encode(b)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}
