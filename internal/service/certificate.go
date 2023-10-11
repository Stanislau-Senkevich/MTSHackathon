package service

import (
	"MTSHackathonBackEnd/internal/entity"
	"MTSHackathonBackEnd/internal/repository"
)

type CertificateService struct {
	repo repository.Certificate
}

func NewCertificateService(repos repository.Certificate) *CertificateService {
	return &CertificateService{repo: repos}
}

func (s *CertificateService) GetAllBoughtCertificates(userId string) ([]entity.Certificate, error) {
	return s.repo.GetAllBoughtCertificates(userId)
}

func (s *CertificateService) ChangeOwnerOfCertificate(id, newUserId string) error {
	return s.repo.ChangeOwnerOfCertificate(id, newUserId)
}

func (s *CertificateService) GenerateLink(id string) string {
	return s.repo.GenerateLink(id)
}

func (s *CertificateService) CreateNewCertificate(cert entity.Certificate) error {
	return s.repo.CreateNewCertificate(cert)
}
