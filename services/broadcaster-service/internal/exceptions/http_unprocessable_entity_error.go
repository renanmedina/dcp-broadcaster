package exceptions

import "net/http"

type HttpUnprocessableEntityError struct {
	msg string
}

func (e HttpUnprocessableEntityError) StatusCode() int {
	return http.StatusUnprocessableEntity
}

func (e HttpUnprocessableEntityError) Error() string {
	return e.msg
}
