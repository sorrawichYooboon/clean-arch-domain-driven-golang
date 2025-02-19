package httperror

type HTTPError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *HTTPError) Error() string {
	return e.Message
}

var (
	ErrBadRequest     = HTTPError{Code: "BAD_REQUEST", Message: "Invalid request"}
	ErrUnauthorized   = HTTPError{Code: "UNAUTHORIZED", Message: "Unauthorized"}
	ErrNotFound       = HTTPError{Code: "NOT_FOUND", Message: "Resource not found"}
	ErrInternalServer = HTTPError{Code: "INTERNAL_ERROR", Message: "Internal server error"}
	ErrConflict       = HTTPError{Code: "CONFLICT", Message: "Conflict, resource exists"}
)
