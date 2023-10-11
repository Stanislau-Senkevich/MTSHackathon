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
}

type Certificate interface {
	GetAllBoughtCertificates(userId string) ([]entity.Certificate, error)
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
