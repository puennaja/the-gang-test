package goerrors

type internalErr struct {
	status  int
	code    int
	message string
	errors  []error
}

func newInternalErr(status, code int, message string, err ...error) *internalErr {
	return &internalErr{
		status:  status,
		code:    code,
		message: message,
		errors:  err,
	}
}

func (i *internalErr) SetStatus(status int) {
	i.status = status
}

func (i *internalErr) SetCode(code int) {
	i.code = code
}

func (i *internalErr) SetMessage(message string) {
	i.message = message
}

func (i *internalErr) SetError(err ...error) {
	i.errors = append(i.errors, err...)
}

func (i *internalErr) GetStatus() int {
	return i.status
}

func (i *internalErr) GetCode() int {
	return i.code
}

func (i *internalErr) GetMessage() string {
	return i.message
}

func (i *internalErr) GetErrors() []error {
	return i.errors
}

func (i *internalErr) Error() string {
	if len(i.errors) == 0 {
		return ""
	}
	var str string
	for _, err := range i.errors {
		str += err.Error() + "\n"
	}
	return str
}
