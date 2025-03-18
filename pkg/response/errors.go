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
	ErrInvalidParams = NewResponse("Invalid parameters", StatusBadRequest, CodeInvalidParams)
	ErrUnauthorized  = NewResponse("Unauthorized access", StatusUnauthorized, CodeUnauthorized)
	ErrForbidden     = NewResponse("Access forbidden", StatusForbidden, CodeForbidden)
	ErrNotFound      = NewResponse("Resource not found", StatusNotFound, CodeNotFound)
	ErrInternal      = NewResponse("Internal server error", StatusInternalError, CodeInternalError)
	ErrDatabase      = NewResponse("Database error", StatusInternalError, CodeDatabaseError)
	ErrDuplicate     = NewResponse("Duplicate entry", StatusConflict, CodeDuplicateEntry)
	ErrValidation    = NewResponse("Validation failed", StatusBadRequest, CodeValidationError)
	ErrService       = NewResponse("Service unavailable", StatusServiceUnavailable, CodeServiceUnavailable)
)

// Success response
var SuccessResponse = NewResponse("Success", StatusOK, CodeSuccess)
