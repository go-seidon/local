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

func NewNotFoundHandler(log logging.Logger, s serialization.Serializer) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Debug("In function: NotFoundHandler")
		defer log.Debug("Returning function: NotFoundHandler")

		w.Header().Set("Content-Type", "application/json")
		Response(WithWriterSerializer(w, s), NotFound("resource not found"))
	}
}

func NewMethodNotAllowedHandler(log logging.Logger, s serialization.Serializer) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Debug("In function: MethodNotAllowedHandler")
		defer log.Debug("Returning function: MethodNotAllowedHandler")

		w.Header().Set("Content-Type", "application/json")
		Response(
			WithWriterSerializer(w, s),
			WithHttpCode(http.StatusMethodNotAllowed),
			Error("method is not allowed"),
		)
	}
}

func NewRootHandler(log logging.Logger, s serialization.Serializer, config *RestAppConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Debug("In function: RootHandler")
		defer log.Debug("Returning function: RootHandler")

		d := struct {
			AppName    string `json:"app_name"`
			AppVersion string `json:"app_version"`
		}{
			AppName:    config.AppName,
			AppVersion: config.AppVersion,
		}
		Response(WithWriterSerializer(w, s), Success(d))
	}
}

func NewHealthCheckHandler(log logging.Logger, s serialization.Serializer, healthService healthcheck.HealthCheck) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Debug("In function: HealthCheckHandler")
		defer log.Debug("Returning function: HealthCheckHandler")

		r, err := healthService.Check()
		if err != nil {

			Response(WithWriterSerializer(w, s), Error(err.Error()))
			return
		}

		jobs := map[string]struct {
			Name      string      `json:"name"`
			Status    string      `json:"status"`
			CheckedAt time.Time   `json:"checked_at"`
			Error     string      `json:"error"`
			Metadata  interface{} `json:"metadata"`
		}{}
		for jobName, item := range r.Items {
			jobs[jobName] = struct {
				Name      string      `json:"name"`
				Status    string      `json:"status"`
				CheckedAt time.Time   `json:"checked_at"`
				Error     string      `json:"error"`
				Metadata  interface{} `json:"metadata"`
			}{
				Name:      item.Name,
				Status:    item.Status,
				Error:     item.Error,
				Metadata:  item.Metadata,
				CheckedAt: item.CheckedAt,
			}
		}

		d := struct {
			Status  string `json:"status"`
			Details map[string]struct {
				Name      string      `json:"name"`
				Status    string      `json:"status"`
				CheckedAt time.Time   `json:"checked_at"`
				Error     string      `json:"error"`
				Metadata  interface{} `json:"metadata"`
			} `json:"details"`
		}{
			Status:  r.Status,
			Details: jobs,
		}

		Response(
			WithWriterSerializer(w, s),
			Success(d),
			WithMessage("success check service health"),
		)
	}
}

func NewDeleteFileHandler(log logging.Logger, s serialization.Serializer, deleter deleting.Deleter) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Debug("In function: DeleteFileHandler")
		defer log.Debug("Returning function: DeleteFileHandler")

		vars := mux.Vars(req)

		ctx := context.Background()
		r, err := deleter.DeleteFile(ctx, deleting.DeleteFileParam{
			FileId: vars["unique_id"],
		})
		if err == nil {

			d := struct {
				DeletedAt int64 `json:"deleted_at"`
			}{
				DeletedAt: r.DeletedAt.UnixMilli(),
			}

			Response(
				WithWriterSerializer(w, s),
				Success(d),
				WithMessage("success delete file"),
			)
			return
		}

		if errors.Is(err, deleting.ErrorResourceNotFound) {
			Response(WithWriterSerializer(w, s), NotFound(err.Error()))
			return
		}

		Response(WithWriterSerializer(w, s), Error(err.Error()))
	}
}

func NewRetrieveFileHandler(log logging.Logger, s serialization.Serializer, retriever retrieving.Retriever) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Debug("In function: RetrieveFileHandler")
		defer log.Debug("Returning function: RetrieveFileHandler")

		vars := mux.Vars(req)

		ctx := context.Background()
		r, err := retriever.RetrieveFile(ctx, retrieving.RetrieveFileParam{
			FileId: vars["unique_id"],
		})
		if err == nil {

			defer r.Data.Close()
			data, err := io.ReadAll(r.Data)
			if err != nil {
				Response(WithWriterSerializer(w, s), Error(err.Error()))
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
			Response(WithWriterSerializer(w, s), NotFound(err.Error()))
			return
		}

		Response(WithWriterSerializer(w, s), Error(err.Error()))
	}
}

func NewUploadFileHandler(log logging.Logger, s serialization.Serializer, uploader uploading.Uploader, locator uploading.UploadLocation, config *RestAppConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Debug("In function: UploadFileHandler")
		defer log.Debug("Returning function: UploadFileHandler")

		// set form max size + add 1KB (non file size estimation if any)
		req.Body = http.MaxBytesReader(w, req.Body, config.UploadFormSize+1024)

		file, fileHeader, err := req.FormFile("file")
		if err != nil {
			Response(WithWriterSerializer(w, s), Error(err.Error()))
			return
		}
		defer file.Close()

		fileInfo, err := ParseMultipartFile(file, fileHeader)
		if err != nil {
			Response(WithWriterSerializer(w, s), Error(err.Error()))
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
			Response(WithWriterSerializer(w, s), Error(err.Error()))
			return
		}

		d := struct {
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
		}

		Response(
			WithWriterSerializer(w, s),
			Success(d),
			WithMessage("success upload file"),
		)
	}
}
