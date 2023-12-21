package ports

import domain "upload_server/internal/core/domain/resource"

type UploadRepository interface {
	HealthCheck() error
	UploadFile(file domain.File) error
}
