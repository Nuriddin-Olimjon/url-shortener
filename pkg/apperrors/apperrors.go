package apperrors

import (
	"errors"
	"fmt"
	"net/http"
)

// Type holds a type string and integer code for the error
type Type string

// "Set" of valid errorTypes
const (
	Authorization        Type = "AUTHORIZATION"        // Authentication Failures -
	Forbidded            Type = "FORBIDDEN"            // Forbidden for requested object
	BadRequest           Type = "BADREQUEST"           // Validation errors / BadInput
	Conflict             Type = "CONFLICT"             // Already exists (eg, create account with existent email) - 409
	Internal             Type = "INTERNAL"             // Server (500) and fallback errors
	NotFound             Type = "NOTFOUND"             // For not finding resource
	PayloadTooLarge      Type = "PAYLOADTOOLARGE"      // for uploading tons of JSON, or an image over the limit - 413
	UnsupportedMediaType Type = "UNSUPPORTEDMEDIATYPE" // for http 415
)

// Error holds a custom error for the application
// which is helpful in returning a consistent
// error type/message from API endpoints
type ErrorResp struct {
	Type        Type   `json:"type"`
	Message     string `json:"message"`
	Code        string `json:"code"`
	InvalidArgs any    `json:"invalid_args,omitempty"`
}

// Error satisfies standard error interface
// we can return errors from this package as
// a regular old go _error_
func (e *ErrorResp) Error() string {
	return e.Message
}

// Status is a mapping errors to status codes
// Of course, this is somewhat redundant since
// our errors already map http status codes
func (e *ErrorResp) Status() int {
	switch e.Type {
	case Authorization:
		return http.StatusUnauthorized
	case BadRequest:
		return http.StatusBadRequest
	case Forbidded:
		return http.StatusForbidden
	case Conflict:
		return http.StatusConflict
	case Internal:
		return http.StatusInternalServerError
	case NotFound:
		return http.StatusNotFound
	case PayloadTooLarge:
		return http.StatusRequestEntityTooLarge
	case UnsupportedMediaType:
		return http.StatusUnsupportedMediaType
	default:
		return http.StatusInternalServerError
	}
}

// GetStatus checks the runtime type
// of the error and returns an http
// status code if the error is model.Error
func GetStatus(err error) int {
	var e *ErrorResp
	if errors.As(err, &e) {
		return e.Status()
	}
	return http.StatusInternalServerError
}

/*
* Error "Factories"
 */

// NewAuthorization to create a 401
func NewAuthorization(reason string) *ErrorResp {
	return &ErrorResp{
		Type:    Authorization,
		Message: reason,
		Code:    "401",
	}
}

func NewForbidden(reason string) *ErrorResp {
	return &ErrorResp{
		Type:    Forbidded,
		Message: reason,
		Code:    "403",
	}
}

// NewBadRequest to create 400 errors (validation, for example)
func NewBadRequest(reason string, args any) *ErrorResp {
	return &ErrorResp{
		Type:        BadRequest,
		Message:     reason,
		Code:        "400",
		InvalidArgs: args,
	}
}

// NewConflict to create an error for 409
func NewConflict(name string, value string) *ErrorResp {
	return &ErrorResp{
		Type:    Conflict,
		Message: fmt.Sprintf("%v with %v already exists", name, value),
		Code:    "409",
	}
}

// NewInternal for 500 errors and unknown errors
func NewInternal(reason string) *ErrorResp {
	return &ErrorResp{
		Type:    Internal,
		Message: fmt.Sprintf("Internal server error. %s", reason),
		Code:    "500",
	}
}

// NewNotFound to create an error for 404
func NewNotFound(name string, value string) *ErrorResp {
	return &ErrorResp{
		Type:    NotFound,
		Message: fmt.Sprintf("%v with %v not found", name, value),
		Code:    "404",
	}
}

// NewPayloadTooLarge to create an error for 413
func NewPayloadTooLarge(maxBodySize int64, contentLength int64) *ErrorResp {
	return &ErrorResp{
		Type:    PayloadTooLarge,
		Message: fmt.Sprintf("Max payload size of %v exceeded. Actual payload size: %v", maxBodySize, contentLength),
		Code:    "413",
	}
}

// NewUnsupportedMediaType to create an error for 415
func NewUnsupportedMediaType(reason string) *ErrorResp {
	return &ErrorResp{
		Type:    UnsupportedMediaType,
		Message: reason,
		Code:    "415",
	}
}
