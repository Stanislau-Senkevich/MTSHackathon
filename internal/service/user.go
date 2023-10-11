package service

import "MTSHackathonBackEnd/internal/repository"

type UserService struct {
	repo repository.User
}

func NewUserService(repos repository.User) *UserService {
	return &UserService{repo: repos}
}
