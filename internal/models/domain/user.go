package domain

type User struct {
	Username string `json:"name" validate:"required"`
}
