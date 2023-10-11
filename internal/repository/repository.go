package repository

import (
	"MTSHackathonBackEnd/internal/entity"
	mongodb "MTSHackathonBackEnd/internal/repository/mongo"
)

type Repository struct {
	User
	Certificate
	Category
	Service
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

func NewRepository(repository *mongodb.MongoRepository) *Repository {
	return &Repository{
		User:        repository,
		Certificate: repository,
		Category:    repository,
	}
}
