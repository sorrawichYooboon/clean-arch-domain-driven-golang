package firebase

type FirebaseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *FirebaseError) Error() string {
	return e.Message
}

var (
	ErrFirebaseUserNotFound = FirebaseError{Code: "ERROR_Firebase_USER_NOT_FOUND", Message: "User not found"}
	ErrFirebaseServiceFail  = FirebaseError{Code: "ERROR_Firebase_SERVICE_FAIL", Message: "Failed to get Firebase service"}
)
