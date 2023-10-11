package service

import (
	"MTSHackathonBackEnd/internal/entity"
	"MTSHackathonBackEnd/internal/repository"
)

type CategoryService struct {
	repo repository.Category
}

func NewCategoryService(repos repository.Category) *CategoryService {
	return &CategoryService{repo: repos}
}

func (s *CategoryService) GetAllCategories() ([]entity.Category, error) {
	return s.repo.GetAllCategories()
}
