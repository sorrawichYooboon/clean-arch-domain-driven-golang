package domain

type Author struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Bio  string `json:"bio"`
}
