package errs

type ErrorCode string

const (
	NotFound     ErrorCode = "NOT_FOUND"
	Conflict     ErrorCode = "CONFLICT"
	Unauthorized ErrorCode = "UNAUTHORIZED"
	Forbidden    ErrorCode = "FORBIDDEN"
	Validation   ErrorCode = "VALIDATION"
	Internal     ErrorCode = "INTERNAL"
)
