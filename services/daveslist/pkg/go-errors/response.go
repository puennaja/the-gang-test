package goerrors

type ResponseErr struct {
	status  int
	Code    int      `json:"code,omitempty"`
	Message string   `json:"message,omitempty"`
	Errors  []string `json:"errors,omitempty"`
}

func newResponseErr(cfg ConfigIFace, err error) *ResponseErr {
	resErr := &ResponseErr{}
	resErr.SetStatus(cfg.DefaultStatus())
	resErr.SetCode(cfg.DefaultCode())
	resErr.SetMessage(cfg.DefaultMessage())

	// return default response error
	if err == nil {
		return resErr
	}

	// check error is internal
	internalErr := IsInternalErr(err)
	if internalErr != nil {
		resErr.SetStatus(internalErr.GetStatus())
		resErr.SetCode(internalErr.GetCode())
		resErr.SetMessage(internalErr.GetMessage())
		resErr.SetError(internalErr.GetErrors()...)
		return resErr
	}

	// only set error
	resErr.SetError(err)
	return resErr
}

func (resErr *ResponseErr) SetStatus(status int) {
	resErr.status = status
}

func (resErr *ResponseErr) SetCode(code int) {
	resErr.Code = code
}

func (resErr *ResponseErr) SetMessage(message string) {
	resErr.Message = message
}

func (resErr *ResponseErr) SetError(err ...error) {
	for _, er := range err {
		resErr.Errors = append(resErr.Errors, er.Error())
	}
}

func (resErr *ResponseErr) GetStatus() int {
	return resErr.status
}

func (resErr *ResponseErr) GetCode() int {
	return resErr.Code
}

func (resErr *ResponseErr) GetMessage() string {
	return resErr.Message
}

func (resErr *ResponseErr) GetErrors() []string {
	return resErr.Errors
}

// error
func (resErr *ResponseErr) Error() string {
	if len(resErr.Errors) == 0 {
		return ""
	}
	var str string
	for _, err := range resErr.Errors {
		str += err + "\n"
	}
	return str
}
