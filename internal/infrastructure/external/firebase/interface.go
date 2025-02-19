package firebase

type FirebaseService interface {
	GetUserByPhoneNumber(request *GetUserByPhoneNumberRequest) (*GetUserByPhoneNumberResponse, error)
}
