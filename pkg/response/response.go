package response

import (
	"github.com/open-runtimes/types-for-go/v4/openruntimes"
)

type Response struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
}

func NewResponse(message string, code int, data interface{}) *Response {
	return &Response{
		Message: message,
		Code:    code,
		Data:    data,
	}
}

func NewJsonResponse(Context *openruntimes.Context, data interface{}) openruntimes.Response {
	response := NewResponse("Success", CodeSuccess, data)
	return Context.Res.Json(response, Context.Res.WithStatusCode(StatusOK))
}

func NewStatusErrorResponse(Context *openruntimes.Context, StatusCode int) openruntimes.Response {
	return Context.Res.Json(map[string]interface{}{}, Context.Res.WithStatusCode(StatusCode))
}

func NewCodeErrorResponse(Context *openruntimes.Context, Response *Response) openruntimes.Response {
	return Context.Res.Json(Response, Context.Res.WithStatusCode(StatusOK))
}
