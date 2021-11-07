package models

type User struct {
	Name     string `json:"Name" validate:"required,min=4,max=16"`
	Email    string `json:"Email" validate:"required,email,min=6,max=32"`
	Password string `json:"Password" validate:"required,min=6,max=16"`
}
