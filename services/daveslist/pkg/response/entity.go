package response

const (
	SuccessCode int = 1000
)

var message = map[int]string{
	SuccessCode: "success",
}

var (
	SuccessResponse = NewResponse(SuccessCode, message[SuccessCode])
)
