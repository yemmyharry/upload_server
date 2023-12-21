package resource

import (
	domain "upload_server/internal/core/domain/resource"
	"upload_server/internal/core/logger"
)

func (m UploadRepositoryDB) HealthCheck() error {
	return nil
}

func (m UploadRepositoryDB) UploadFile(file domain.File) error {
	result := m.db.Where("name = ?", file.Name).First(&domain.File{})
	if result.Error == nil {
		logger.Info("File already exists")
		return result.Error
	}
	m.db.Create(&file)
	return nil
}
