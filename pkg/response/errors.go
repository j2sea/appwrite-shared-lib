package response

// HTTP Status codes
const (
	StatusOK                 = 200
	StatusBadRequest         = 400
	StatusUnauthorized       = 401
	StatusForbidden          = 403
	StatusNotFound           = 404
	StatusMethodNotAllowed   = 405
	StatusConflict           = 409
	StatusInternalError      = 500
	StatusServiceUnavailable = 503
)

// Error codes
const (
	CodeSuccess            = 0
	CodeInvalidParams      = 1001
	CodeUnauthorized       = 1002
	CodeForbidden          = 1003
	CodeNotFound           = 1004
	CodeInternalError      = 1005
	CodeDatabaseError      = 1006
	CodeDuplicateEntry     = 1007
	CodeValidationError    = 1008
	CodeServiceUnavailable = 1009
)

// Predefined error responses
var (
	ErrInvalidParams = NewResponse("Invalid parameters", CodeInvalidParams, nil)
	ErrUnauthorized  = NewResponse("Unauthorized access", CodeUnauthorized, nil)
	ErrForbidden     = NewResponse("Access forbidden", CodeForbidden, nil)
	ErrNotFound      = NewResponse("Resource not found", CodeNotFound, nil)
	ErrInternal      = NewResponse("Internal server error", CodeInternalError, nil)
	ErrDatabase      = NewResponse("Database error", CodeDatabaseError, nil)
	ErrDuplicate     = NewResponse("Duplicate entry", CodeDuplicateEntry, nil)
	ErrValidation    = NewResponse("Validation failed", CodeValidationError, nil)
	ErrService       = NewResponse("Service unavailable", CodeServiceUnavailable, nil)
)
