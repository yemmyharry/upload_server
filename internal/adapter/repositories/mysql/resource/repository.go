package resource

import (
	"os"
)

type UploadRepositoryDB struct{}

func (m UploadRepositoryDB) HealthCheck() error {
	return nil
}

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
