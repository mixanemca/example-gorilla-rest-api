// Package service for behavior & logic
package service

import (
	"net/http"

	v1 "github.com/mixanemca/example-gorilla-rest-api/internal/app/api/handler/v1"
)

type Service struct {
	userRepo v1.UserRepository
}

func NewService(userRepo v1.UserRepository) *Service {
	return &Service{
		userRepo: userRepo,
	}
}

func (s Service) CreateUser(w http.ResponseWriter, r *http.Request) {
	s.userRepo.CreateUser(w, r)
}

func (s Service) GetUserByID(w http.ResponseWriter, r *http.Request) {
	s.userRepo.GetUserByID(w, r)
}

func (s Service) GetUsers(w http.ResponseWriter, r *http.Request) {
	s.userRepo.GetUsers(w, r)
}

func (s Service) UpdateUser(w http.ResponseWriter, r *http.Request) {
	s.userRepo.UpdateUser(w, r)
}

func (s Service) DeleteUser(w http.ResponseWriter, r *http.Request) {
	s.userRepo.DeleteUser(w, r)
}
