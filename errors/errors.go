package errors

type Error interface {
	error
	GetCode() int64
}

type StatusError struct {
	Code    int64         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func (e StatusError) Error() string {
	return e.Message
}

func (e StatusError) GetCode() int64 {
	return e.Code
}

func New(code int64, message string) Error {
	return StatusError{
		code,
		message,
		nil,
	}
}
