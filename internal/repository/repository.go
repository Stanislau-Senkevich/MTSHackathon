package repository

import (
	mongodb "MTSHackathonBackEnd/internal/repository/mongo"
)

type Repository struct {
	User
	Certificate
	Category
}

type User interface {
}

type Certificate interface {
}

type Category interface {
}

func NewRepository(repository *mongodb.MongoRepository) *Repository {
	return &Repository{
		User:        repository,
		Certificate: repository,
		Category:    repository,
	}
}
