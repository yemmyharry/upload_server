package ports

import domain "upload_server/internal/core/domain/resource"

type UploadService interface {
	HealthCheck() error
	UploadFile(file domain.File) error
	GetFiles() ([]domain.File, error)
}
