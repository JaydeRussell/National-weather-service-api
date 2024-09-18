package data

import "fmt"

type HTTPError struct {
	StatusCode int
	Message    string
}

func NewHttpError(message string, code int) *HTTPError {
	return &HTTPError{
		Message:    message,
		StatusCode: code,
	}
}

func (h HTTPError) Error() string {
	return fmt.Sprintf("{error: %s, code: %d}", h.Message, h.StatusCode)
}
