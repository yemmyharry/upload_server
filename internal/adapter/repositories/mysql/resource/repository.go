package resource

import (
	"os"
)

func (m UploadRepositoryDB) HealthCheck() error {
	return nil
}

//func (m UploadRepositoryDB) UploadFile(file domain.File) error {
//	result := m.db.Where("name = ?", file.Name).First(&domain.File{})
//	if result.Error == nil {
//		logger.Info("File already exists")
//		return result.Error
//	}
//	m.db.Create(&file)
//	return nil
//}

func (m UploadRepositoryDB) GetFiles() ([]string, error) {

	var fileNames []string

	files, err := os.ReadDir(os.Getenv("DIRECTORY"))
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	return fileNames, nil
}
