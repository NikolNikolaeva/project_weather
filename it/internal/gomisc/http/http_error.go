package http

type APIError interface {
	error

	Code() int // see: https://go.dev/src/net/http/status.go
	Body() any
	Cause() error
}

func WrapError(cause error, code int, body any) APIError {
	if apiError, ok := cause.(APIError); ok {
		return apiError
	}

	return NewAPIError(cause, code, body)
}

func NewAPIError(cause error, code int, body any) APIError {
	if cause == nil {
		return nil
	}

	return &_APIError{
		body:  body,
		code:  code,
		cause: cause,
	}
}

type _APIError struct {
	body  any
	code  int
	cause error
}

func (self *_APIError) Error() string {
	return self.cause.Error()
}

func (self *_APIError) Code() int {
	return self.code
}

func (self *_APIError) Body() any {
	return self.body
}

func (self *_APIError) Cause() error {
	return self.cause
}
