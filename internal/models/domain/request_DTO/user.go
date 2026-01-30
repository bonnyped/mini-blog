package domain

type User struct {
	Username string `json:"username" validate:"required,min=3"`
}
