package errs

import (
	"errors"
	"fmt"
)

// ErrorCode represents a domain-level semantic error classification.
// These codes are protocol-agnostic — adapters are responsible for
// translating them into transport-specific responses (HTTP, gRPC, etc.).
type ErrorCode string

const (
	ErrNotFound     ErrorCode = "NOT_FOUND"
	ErrConflict     ErrorCode = "CONFLICT"
	ErrUnauthorized ErrorCode = "UNAUTHORIZED"
	ErrForbidden    ErrorCode = "FORBIDDEN"
	ErrValidation   ErrorCode = "VALIDATION"
	ErrInternal     ErrorCode = "INTERNAL"
)

// DomainError is a domain-level error with a semantic code and message.
type DomainError struct {
	Code    ErrorCode
	Message string
	Err     error // wrapped original error
}

// Error implements the error interface.
func (e *DomainError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%s] %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

// Unwrap returns the wrapped error for errors.Is/As chain support.
func (e *DomainError) Unwrap() error {
	return e.Err
}

// NewNotFound creates a NOT_FOUND domain error.
func NewNotFound(message string, args ...any) *DomainError {
	return &DomainError{
		Code:    ErrNotFound,
		Message: fmt.Sprintf(message, args...),
	}
}

// NewConflict creates a CONFLICT domain error.
func NewConflict(message string, args ...any) *DomainError {
	return &DomainError{
		Code:    ErrConflict,
		Message: fmt.Sprintf(message, args...),
	}
}

// NewUnauthorized creates an UNAUTHORIZED domain error.
func NewUnauthorized(message string, args ...any) *DomainError {
	return &DomainError{
		Code:    ErrUnauthorized,
		Message: fmt.Sprintf(message, args...),
	}
}

// NewForbidden creates a FORBIDDEN domain error.
func NewForbidden(message string, args ...any) *DomainError {
	return &DomainError{
		Code:    ErrForbidden,
		Message: fmt.Sprintf(message, args...),
	}
}

// NewValidation creates a VALIDATION domain error.
func NewValidation(message string, args ...any) *DomainError {
	return &DomainError{
		Code:    ErrValidation,
		Message: fmt.Sprintf(message, args...),
	}
}

// NewInternal creates an INTERNAL domain error wrapping the original error.
func NewInternal(err error, message string, args ...any) *DomainError {
	return &DomainError{
		Code:    ErrInternal,
		Message: fmt.Sprintf(message, args...),
		Err:     err,
	}
}

// GetCode extracts the ErrorCode from any error.
// Returns ErrInternal if the error is not a DomainError.
func GetCode(err error) ErrorCode {
	if de, ok := errors.AsType[*DomainError](err); ok {
		return de.Code
	}
	return ErrInternal
}

// IsDomainError reports whether err is of type *DomainError or wraps one.
func IsDomainError(err error) bool {
	var de *DomainError
	return errors.As(err, &de)
}

// Domain error messages — shared constants for auth-related errors.
const (
	ErrMsgEmailAlreadyExists = "email already exists"
	ErrMsgUserNotFound       = "user not found"
	ErrMsgInvalidCredentials = "invalid email or password"
	ErrMsgInvalidToken       = "invalid or expired token"
	ErrMsgTokenMissing       = "token is missing"
	ErrMsgUnauthorized       = "unauthorized access"
)
