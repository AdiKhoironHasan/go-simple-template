package errs

import "errors"

const (
	ErrMsgEmailAlreadyExists = "email already exists"
	ErrMsgUserNotFound       = "user not found"
	ErrMsgInvalidCredentials = "invalid email or password"
	ErrMsgInvalidToken       = "invalid or expired token"
	ErrMsgTokenMissing       = "token is missing"
	ErrMsgUnauthorized       = "unauthorized access"
)

// Auth errors
var (
	ErrInvalidCredentials = errors.New(ErrMsgInvalidCredentials)
	ErrEmailAlreadyExists = errors.New(ErrMsgEmailAlreadyExists)
	ErrUserNotFound       = errors.New(ErrMsgUserNotFound)
	ErrInvalidToken       = errors.New(ErrMsgInvalidToken)
	ErrTokenMissing       = errors.New(ErrMsgTokenMissing)
	ErrUnauthorized       = errors.New(ErrMsgUnauthorized)
)
