package apperror

type UseCaseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *UseCaseError) Error() string {
	return e.Message
}

var (
	ErrUserExists         = UseCaseError{Code: "ERROR_USER_EXISTS", Message: "User already exists"}
	ErrUserNotFound       = UseCaseError{Code: "ERROR_USER_NOT_FOUND", Message: "User not found"}
	ErrEmailConflict      = UseCaseError{Code: "ERROR_EMAIL_CONFLICT", Message: "Email is already associated with another account"}
	ErrInvalidInput       = UseCaseError{Code: "ERROR_INVALID_INPUT", Message: "Invalid input data"}
	ErrDatabase           = UseCaseError{Code: "ERROR_DATABASE", Message: "Database error"}
	ErrCacheDatabase      = UseCaseError{Code: "ERROR_CACHE_DATABASE", Message: "Cache database error"}
	ErrHashPassword       = UseCaseError{Code: "ERROR_HASH_PASSWORD", Message: "Error hashing password"}
	ErrInvalidCredentials = UseCaseError{Code: "ERROR_INVALID_CREDENTIALS", Message: "Invalid credentials"}
)
