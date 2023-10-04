// Package service for behavior & logic
package service

import v1 "github.com/mixanemca/example-gorilla-rest-api/internal/app/api/handler/v1"

type Service struct {
	userRepo v1.UserRepository
}

func NewService(userRepo v1.UserRepository) *Service {
	return &Service{
		userRepo: userRepo,
	}
}
