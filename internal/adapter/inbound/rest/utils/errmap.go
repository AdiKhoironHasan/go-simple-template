package utils

import (
	"net/http"

	"github.com/adikhoironhasan/go-simple-template/internal/pkg/errs"
)

// codeToHTTP maps domain error codes to HTTP status codes.
var codeToHTTP = map[errs.ErrorCode]int{
	errs.ErrNotFound:     http.StatusNotFound,
	errs.ErrConflict:     http.StatusConflict,
	errs.ErrUnauthorized: http.StatusUnauthorized,
	errs.ErrForbidden:    http.StatusForbidden,
	errs.ErrValidation:   http.StatusUnprocessableEntity,
	errs.ErrInternal:     http.StatusInternalServerError,
}

// MapErrorToHTTP translates a domain error into an HTTP status code.
// Returns 500 Internal Server Error for unknown or non-domain errors.
func MapErrorToHTTP(err error) int {
	code := errs.GetCode(err)
	if status, ok := codeToHTTP[code]; ok {
		return status
	}
	return http.StatusInternalServerError
}
