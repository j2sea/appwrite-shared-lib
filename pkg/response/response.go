package response

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
