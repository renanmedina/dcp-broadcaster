package exceptions

import "net/http"

type HttpResponseError interface {
	StatusCode() int
	Error() string
}

type DefaultHttpResponseError struct {
	statusCode int
	msg        string
}

func (e DefaultHttpResponseError) StatusCode() int {
	return http.StatusUnprocessableEntity
}

func (e DefaultHttpResponseError) Error() string {
	return e.msg
}

func NewHttpResponseError(code int, message string) HttpResponseError {
	if code == http.StatusUnprocessableEntity {
		return HttpUnprocessableEntityError{message}
	}

	return DefaultHttpResponseError{code, message}
}
