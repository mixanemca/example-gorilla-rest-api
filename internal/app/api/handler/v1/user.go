// Package v1 for endpoints
package v1

import (
	"net/http"
)

//go:generate mockgen -source=user.go -destination=mocks/mock.go

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

// @Summary Create user
// @Tags users
// @Description Create user
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.IDInfo
// @Failure 400
// @Failure 403
// @Failure 500
// @router /user [post]
func (u userService) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("not implemented"))
}
