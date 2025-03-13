package error_custom

import (
	"fmt"
	"net/http"
)

type ErrorModel struct {
	StatusCode int
	Message    string
}

func (e *ErrorModel) Error() string {
	return fmt.Sprintf("Status: %d, Message: %s", e.StatusCode, e.Message)
}

func NewBadRequest(message string) ErrorModel {
	return ErrorModel{
		StatusCode: http.StatusBadRequest,
		Message:    message,
	}
}

func NewInternalServerError(message string) ErrorModel {
	return ErrorModel{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
	}
}
