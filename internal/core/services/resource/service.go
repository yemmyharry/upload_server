package services

import (
	ports "upload_server/internal/ports/resource"
)

type Service struct {
	uploadRepository ports.UploadRepository
}

func New(uploadRepository ports.UploadRepository) *Service {
	return &Service{
		uploadRepository: uploadRepository,
	}
}

func (s Service) HealthCheck() error {
	return s.uploadRepository.HealthCheck()
}
