package dto

// Response represents a standard API response structure.
// @Description Standard response structure for API responses.
// @Model
type ResponseDTO struct {
	Message string `json:"message" example:"User created successfully"`
}
