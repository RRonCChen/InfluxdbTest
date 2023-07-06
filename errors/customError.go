package errors

type CustomError struct {
	Code    int
	Message string
}

var (
	InvaildParamError = NewError(400, "Invaild params")
	NotFoundError     = NewError(404, "Data not found")
)

func (err *CustomError) Error() string {
	return err.Message
}

func NewError(code int, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}

func (err *CustomError) SetMessage(message string) CustomError {
	err.Message = message
	return *err
}
