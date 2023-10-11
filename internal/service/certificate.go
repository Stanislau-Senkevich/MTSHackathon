package service

import "MTSHackathonBackEnd/internal/repository"

type CertificateService struct {
	repo repository.Certificate
}

func NewCertificateService(repos repository.Certificate) *CertificateService {
	return &CertificateService{repo: repos}
}
