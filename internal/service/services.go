package service

import "MTSHackathonBackEnd/internal/repository"

type Services struct {
	User
	Certificate
	Category
}

type Dependencies struct {
}

type User interface {
}

type Certificate interface {
}

type Category interface {
}

func NewService(repos *repository.Repository) *Services {
	return &Services{
		User:        NewUserService(repos),
		Certificate: NewCertificateService(repos),
		Category:    NewCategoryService(repos),
	}
}
