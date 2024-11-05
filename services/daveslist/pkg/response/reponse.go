package response

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponse(code int, message string) *Response {
	return &Response{
		Code:    code,
		Message: message,
	}
}

func (r *Response) SetData(data interface{}) *Response {
	r.Data = data
	return r
}

func (r *Response) SetCode(code int) *Response {
	r.Code = code
	return r
}

func (r *Response) SetMessage(message string) *Response {
	r.Message = message
	return r
}
