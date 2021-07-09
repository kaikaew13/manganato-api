package manganatoapi

import "net/http"

type NotFoundError struct {
	msg        string
	statusCode int
}

func newNotFoundError() *NotFoundError {
	return &NotFoundError{
		msg:        "This page does not exist or has been deleted",
		statusCode: http.StatusNotFound,
	}
}

func (e *NotFoundError) Error() string {
	return e.msg
}

func (e *NotFoundError) getStatusCode() int {
	return e.statusCode
}
