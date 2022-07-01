package rest_app

const (
	CODE_SUCCESS   = "SUCCESS"
	CODE_ERROR     = "ERROR"
	CODE_NOT_FOUND = "NOT_FOUND"
)

type ResponseBody struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type NewResponseBodyParam struct {
	Code    string
	Message string
	Data    interface{}
}

func NewResponseBody(p *NewResponseBodyParam) ResponseBody {
	r := ResponseBody{
		Message: "success",
		Code:    CODE_SUCCESS,
	}
	if p == nil {
		return r
	}
	if p.Code != "" {
		r.Code = p.Code
	}
	if p.Message != "" {
		r.Message = p.Message
	}
	if p.Data != nil {
		r.Data = p.Data
	}
	return r
}
