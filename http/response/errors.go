package response

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrNotFound            = errors.New("not found")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrInternalServerError = errors.New("internal server error")
	ErrUnknownHTTPError    = errors.New("unknown error")
)

// HTTPError represents an http error
type HTTPError struct {
	Status int   `json:"-"`
	Errors any   `json:"errors,omitempty"`
	Causes error `json:"-"`
}

// Error satisfy error interface
func (h HTTPError) Error() string {
	return fmt.Sprintf("status: %d - errors: %s - causes: %v", h.Status, h.Errors, h.Causes)
}

// A NewHTTPError is a factory function for HTTP Errors
func NewHTTPError(status int, err string, causes error) HTTPError {
	return HTTPError{
		Status: status,
		Errors: err,
		Causes: causes,
	}
}

// NewBadRequestError represents an HTTP error
// with status code of 400
func NewBadRequestError(errors any, causes error) HTTPError {
	return HTTPError{
		Status: http.StatusBadRequest,
		Errors: errors,
		Causes: causes,
	}
}

// NewNotFoundError represents an HTTP error
// with status code of 404
func NewNotFoundError(causes error) HTTPError {
	return HTTPError{
		Status: http.StatusNotFound,
		Errors: ErrNotFound.Error(),
		Causes: causes,
	}
}

// NewUnauthorizedError represents an HTTP error
// with status code of 401
func NewUnauthorizedError(causes error) HTTPError {
	return HTTPError{
		Status: http.StatusUnauthorized,
		Errors: ErrUnauthorized.Error(),
		Causes: causes,
	}
}

// NewInternalServerError represents an HTTP error
// with status code of 500
func NewInternalServerError(causes error) HTTPError {
	result := HTTPError{
		Status: http.StatusInternalServerError,
		Errors: ErrInternalServerError.Error(),
		Causes: causes,
	}
	return result
}

// ParseHTTPError parses an error
// to HTTPError, returns 500
// if given error is not instance
// of HTTPError
func ParseHTTPError(err error) HTTPError {
	var httpErr HTTPError
	if errors.As(err, &httpErr) {
		return httpErr
	}
	return NewHTTPError(
		http.StatusInternalServerError,
		ErrUnknownHTTPError.Error(),
		err,
	)
}
