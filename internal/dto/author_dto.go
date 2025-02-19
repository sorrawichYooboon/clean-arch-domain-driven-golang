package dto

type GetAllAuthorResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

type CreateAuthorRequest struct {
	Name string `json:"name" validate:"required"`
	Bio  string `json:"bio" validate:"required"`
}

type UpdateAuthorRequest struct {
	Name string `json:"name" validate:"required"`
	Bio  string `json:"bio" validate:"required"`
}

type UpdateAuthorResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Bio  string `json:"bio"`
}
