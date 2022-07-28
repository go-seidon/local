package rest_app

import (
	"fmt"
	"net/http"

	"github.com/go-seidon/local/internal/serialization"
)

const (
	CODE_SUCCESS      = "SUCCESS"
	CODE_ERROR        = "ERROR"
	CODE_NOT_FOUND    = "NOT_FOUND"
	CODE_UNAUTHORIZED = "UNAUTHORIZED"
)

type ResponseBody struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseParam struct {
	Writer     http.ResponseWriter
	Serializer serialization.Serializer

	Body            ResponseBody
	defaultHttpCode int
	HttpCode        int
	Message         string
	Code            string
}

type ResponseOption = func(*ResponseParam)

func WithWriterSerializer(w http.ResponseWriter, s serialization.Serializer) ResponseOption {
	return func(rp *ResponseParam) {
		rp.Writer = w
		rp.Serializer = s
	}
}

func WithHttpCode(c int) ResponseOption {
	return func(rp *ResponseParam) {
		rp.HttpCode = c
	}
}

func WithMessage(m string) ResponseOption {
	return func(rp *ResponseParam) {
		rp.Message = m
	}
}

func WithCode(c string) ResponseOption {
	return func(rp *ResponseParam) {
		rp.Code = c
	}
}

func Success(d interface{}) ResponseOption {
	return func(rp *ResponseParam) {
		b := ResponseBody{
			Message: "success",
			Code:    CODE_SUCCESS,
			Data:    d,
		}

		rp.Body = b
		rp.defaultHttpCode = http.StatusOK
	}
}

func Error(message string) ResponseOption {
	return func(rp *ResponseParam) {
		b := ResponseBody{
			Message: "error",
			Code:    CODE_ERROR,
		}
		if message != "" {
			b.Message = message
		}

		rp.Body = b
		rp.defaultHttpCode = http.StatusBadRequest
	}
}

func Response(opts ...ResponseOption) error {
	p := ResponseParam{
		Body: ResponseBody{
			Code:    CODE_SUCCESS,
			Message: "success",
		},
		defaultHttpCode: http.StatusOK,
	}
	for _, opt := range opts {
		opt(&p)
	}

	if p.Writer == nil {
		return fmt.Errorf("writer should be specified")
	}
	if p.Serializer == nil {
		return fmt.Errorf("serializer should be specified")
	}

	httpCode := p.defaultHttpCode
	if p.HttpCode != 0 {
		httpCode = p.HttpCode
	}

	if p.Message != "" {
		p.Body.Message = p.Message
	}

	if p.Code != "" {
		p.Body.Code = p.Code
	}

	r, err := p.Serializer.Marshal(p.Body)
	if err != nil {
		return err
	}

	p.Writer.WriteHeader(httpCode)
	p.Writer.Write(r)
	return nil
}
