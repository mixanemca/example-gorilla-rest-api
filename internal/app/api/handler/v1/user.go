// Package v1 for interface UserRepository
package v1

import (
	"net/http"
)

//go:generate mockgen -source=user.go -destination=mocks/mock.go

type UserRepository interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUserByID(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}
