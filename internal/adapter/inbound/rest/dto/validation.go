package dto

import (
	"context"
	"fmt"
	"log/slog"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/adikhoironhasan/go-simple-template/internal/pkg/consts"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// ValidateRequest validates the request and returns the response
func ValidateRequest(ctx context.Context, errParse, errValidate error) (*Response, error) {
	if errParse != nil {
		slog.ErrorContext(ctx, "error parsing request", slog.String(consts.Error, errParse.Error()))
		response := RestResponse(
			http.StatusBadRequest,
			nil,
			errParse,
		)
		return response, fmt.Errorf("invalid request body or parameter: %w", errParse)
	}

	if errValidate != nil {
		response := RestResponse(
			http.StatusUnprocessableEntity,
			nil,
			errValidate,
		)
		return response, errValidate
	}

	return nil, nil
}

// ParseValidationErrors parses the validation errors into a slice of Errors
func ParseValidationErrors(err error, prefix ...string) []Error {
	var errs []Error

	switch ve := err.(type) {
	case validation.Errors:
		errs = parseMultipleValidationErrors(ve, prefix...)
	case validation.Error:
		errs = parseSingleValidationError(ve, prefix...)
	default:
		errs = parseUnknownError(err)
	}

	return errs
}

// parseMultipleValidationErrors handles validation.Errors case
func parseMultipleValidationErrors(ve validation.Errors, prefix ...string) []Error {
	var errs []Error

	for field, validationErr := range ve {
		if validationErr != nil {
			if nestedErr, ok := validationErr.(validation.Errors); ok {
				nestedErrs := ParseValidationErrors(nestedErr, append(prefix, field)...)
				errs = append(errs, nestedErrs...)
				continue
			}

			fullField := buildFullFieldName(field, prefix...)
			errs = append(errs, Error{
				Field:   fullField,
				Message: validationErr.Error(),
			})
		}
	}

	return errs
}

// parseSingleValidationError handles validation.Error case
func parseSingleValidationError(ve validation.Error, prefix ...string) []Error {
	field := buildFullFieldName(ve.Code(), prefix...)
	return []Error{{
		Field:   field,
		Message: ve.Error(),
	}}
}

// parseUnknownError handles unknown error types
func parseUnknownError(err error) []Error {
	return []Error{{
		Message: err.Error(),
	}}
}

// buildFullFieldName constructs the full field name with prefix
func buildFullFieldName(field string, prefix ...string) string {
	if len(prefix) > 0 {
		return fmt.Sprintf("%s.%s", strings.Join(prefix, "."), field)
	}
	return field
}

func validateFile(field string, files []*multipart.FileHeader, exts []string, maxSizeMb int) validation.RuleFunc {
	return func(_ any) error {
		if len(files) == 0 {
			return validation.NewError(field, "file is required")
		}

		for _, file := range files {
			if file.Size > int64(maxSizeMb*1024*1024) {
				return validation.NewError(field, fmt.Sprintf("file size must be less than %d MB", maxSizeMb))
			}

			fileExt := filepath.Ext(file.Filename)
			isAllowExt := false
			for _, ext := range exts {
				if fileExt == ext {
					isAllowExt = true
					break
				}
			}

			if !isAllowExt {
				return validation.NewError(field, fmt.Sprintf("file extension must be %v", exts))
			}

		}

		return nil
	}
}
