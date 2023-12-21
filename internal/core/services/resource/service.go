package services

import (
	domain "upload_server/internal/core/domain/resource"
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

func (s Service) UploadFile(file domain.File) error {
	return s.uploadRepository.UploadFile(file)
}

func (s Service) GetFiles() ([]domain.File, error) {
	return s.uploadRepository.GetFiles()
}

func (s Service) DownloadFile(filename string) (*domain.File, error) {
	return s.uploadRepository.DownloadFile(filename)
}
