// Package v1 for endpoints
package v1

import (
	"net/http"
)

type UserRepository interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}
type userService struct {
	//db *db.conn
}

func NewUserRepository() *userService {
	return &userService{
		//db: db
	}
}

func (u userService) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("not implemented"))
}
