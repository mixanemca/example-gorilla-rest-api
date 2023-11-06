// Package models for users
package models

type Users []UserResponse

type User struct {
	BaseModel
	Name     string `json:"name" validate:"required,alpha,min=3"`
	Surname  string `json:"surname" validate:"required,alpha,min=3"`
	Username string `json:"username" validate:"required,alphanum,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required,e164"`
}

type UserResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

// TODO: convert to entity and UserRequest struct
