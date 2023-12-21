package resource

import (
	"errors"
	"os"
)

type UploadRepositoryDB struct{}

func (m UploadRepositoryDB) HealthCheck() error {
	return nil
}

func (m UploadRepositoryDB) GetFiles() ([]string, error) {

	var fileNames []string

	files, err := os.ReadDir("files")
	if err != nil {
		return nil, errors.New("no file uploaded yet")
	}

	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	return fileNames, nil
}
