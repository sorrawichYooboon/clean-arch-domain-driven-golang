package dto

// Error represents a standard error response.
// @Description Standard error response
// @Model
type Error struct {
	Message string `json:"message" example:"An error occurred"`
}
