package service

import "MTSHackathonBackEnd/internal/repository"

type CategoryService struct {
	repo repository.Category
}

func NewCategoryService(repos repository.Category) *CategoryService {
	return &CategoryService{repo: repos}
}
