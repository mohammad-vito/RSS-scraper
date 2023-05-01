package validate

import (
	"errors"
)

type ErrorResponse struct {
	Error  string       `json:"error"`
	Fields []ErrorField `json:"fields,omitempty"`
}

type ErrorField struct {
	Field string `json:"field"`
	Err   string `json:"error"`
}

type RequestError struct {
	Err    error
	Status int
	Fields error
}

func NewRequestError(err error, status int) error {
	return &RequestError{err, status, nil}
}

func (e *RequestError) Error() string {
	return e.Err.Error()
}

// Cause iterates through all the wrapped errors until the root
// error value is reached.
func Cause(err error) error {
	root := err
	for {
		if err = errors.Unwrap(root); err == nil {
			return root
		}
		root = err
	}
}
