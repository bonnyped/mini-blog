package domain

type User struct {
	Username string `json:"name" validate:"required"`
	Password string `json:"-" validate:"reequired,min=8"`
}
