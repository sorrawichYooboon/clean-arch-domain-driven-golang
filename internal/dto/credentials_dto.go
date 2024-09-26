package dto

// Credentials represents the login credentials for a user.
// @Description User login credentials that include username and password.
// @Model
type CredentialsDTO struct {
	Username string `json:"username" example:"john_doe"`
	Password string `json:"password" example:"password123"`
}

// TokenResponse represents the response containing the JWT token.
// @Description JWT token response after successful login.
// @Model
type TokenResponseDTO struct {
	Token string `json:"token" example:"your_jwt_token_here"`
}
