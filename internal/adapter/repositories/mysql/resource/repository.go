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

func (m UploadRepositoryDB) GetFiles() ([]domain.File, error) {
	var files []domain.File
	m.db.Find(&files)
	return files, nil
}

func (m UploadRepositoryDB) DownloadFile(filename string) (*domain.File, error) {
	var file domain.File
	if err := m.db.Where("name = ?", filename).First(&file).Error; err != nil {
		logger.Error("Error getting file: " + err.Error())
		return nil, err
	}
	return &file, nil
}
