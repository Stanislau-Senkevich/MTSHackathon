package service

import (
	"MTSHackathonBackEnd/internal/entity"
	"MTSHackathonBackEnd/internal/repository"
)

type ServiceService struct {
	repo repository.Service
}

func NewServiceService(repos repository.Service) *ServiceService {
	return &ServiceService{repo: repos}
}

func (s *ServiceService) GetAllServices() ([]entity.Service, error) {
	return s.repo.GetAllServices()
}
