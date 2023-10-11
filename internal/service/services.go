package service

import (
	"MTSHackathonBackEnd/internal/entity"
	"MTSHackathonBackEnd/internal/repository"
)

type Services struct {
	User
	Certificate
	Category
	Service
}

type Dependencies struct {
}

type User interface {
	GetUserIdByPhoneNumber(phoneNumber string) (string, error)
	CheckUser(id string) bool
}

type Certificate interface {
	GetAllBoughtCertificates(userId string) ([]entity.Certificate, error)
	ChangeOwnerOfCertificate(id, newUserId string) error
	GenerateLink(id string) string
	CreateNewCertificate(cert entity.Certificate) error
}

type Category interface {
	GetAllCategories() ([]entity.Category, error)
}

type Service interface {
	GetAllServices() ([]entity.Service, error)
}

func NewService(repos *repository.Repository) *Services {
	return &Services{
		User:        NewUserService(repos),
		Certificate: NewCertificateService(repos),
		Category:    NewCategoryService(repos),
		Service:     NewServiceService(repos),
	}
}
