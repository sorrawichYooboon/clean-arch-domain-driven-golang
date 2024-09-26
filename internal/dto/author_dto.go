package dto

type AuthorDTO struct {
	Name string `json:"name" validate:"required"`
	Bio  string `json:"bio" validate:"required"`
}
