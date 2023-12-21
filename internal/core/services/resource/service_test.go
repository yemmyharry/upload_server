package services

import (
	"errors"
	"github.com/golang/mock/gomock"
	"testing"
	services "upload_server/internal/core/services/mock"
)

func TestApplication_GetFiles(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockedRepository := services.NewMockUploadRepository(ctrl)
	service := New(mockedRepository)
	t.Run("Test for an error", func(t *testing.T) {
		mockedRepository.EXPECT().GetFiles().Return([]string{}, errors.New("an error occurred"))
		_, err := service.GetFiles()
		if err.Error() != "an error occurred" {
			t.Errorf("Expected error: %s, got: %s", "an error occurred", err.Error())
		}
	})
	t.Run("Test for all files", func(t *testing.T) {
		mockedRepository.EXPECT().GetFiles().Return([]string{"test_file_one", "test_file_2"}, nil)

		resp, err := service.GetFiles()
		if err != nil {
			t.Errorf("Expected error: %v, got: %s", nil, err.Error())
		}
		if resp[0] != "test_file_one" {
			t.Errorf("Expected balance: %s, got: %s", "100", resp[0])
		}
	})
}
