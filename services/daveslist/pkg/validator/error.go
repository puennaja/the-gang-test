package validator

type Error struct {
	message string
}

func NewError(text string) Error {
	return Error{message: text}
}

func (e Error) Error() string {
	return e.message
}
