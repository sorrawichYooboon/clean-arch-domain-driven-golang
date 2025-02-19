package dto

// UserLoginRequest represents the login credentials for a user.
// @Description User login credentials that include username and password.
type UserLoginRequest struct {
	Username string `json:"username" example:"john_doe"`
	Password string `json:"password" example:"password123"`
}

// UserRegisterRequest represents the registration data for a new user.
// @Description User registration data.
type UserRegisterRequest struct {
	Username string `json:"username" example:"john_doe"`
	Password string `json:"password" example:"password123"`
}

// UserLoginResponse represents the response containing the JWT token.
// @Description JWT token response after successful login.
type UserLoginResponse struct {
	Token string `json:"token" example:"your_jwt_token_here"`
}
