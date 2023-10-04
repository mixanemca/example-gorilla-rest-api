// Package v1 for endpoints
package v1

import "net/http"

type UserRepository interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}

var Instance UserRepository = userService{
	//db: db.GetDB()
}

type userService struct {
	//db *db.conn
}

func (u userService) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("not implemented"))
}
