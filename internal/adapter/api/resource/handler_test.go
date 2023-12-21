package resource

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	services "upload_server/internal/core/services/mock"
)

func TestApplication_GetAllFiles(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockedService := services.NewMockUploadService(ctrl)
	handler := NewHTTPHandler(mockedService)

	router := gin.Default()

	handler.Routes(router)

	t.Run("Get all files", func(t *testing.T) {

		mockedService.EXPECT().GetFiles().Return([]string{"file_one", "file_2", "file_three"}, nil)
		req, err := http.NewRequest(http.MethodGet, "/api/v1/files", nil)
		if err != nil {
			t.Fatal(err)
		}
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)
		if !strings.Contains(resp.Body.String(), "one") {
			t.Fatalf("Expected balance: %s to be in the response body", "one")
		}
		if !strings.Contains(resp.Body.String(), "2") {
			t.Fatalf("Expected balance: %s to be in the response body", "2")
		}
		if resp.Code != http.StatusOK {
			t.Fatalf("Expected status code: %d but got: %d", http.StatusOK, resp.Code)
		}
	})

	t.Run("should return an error", func(t *testing.T) {
		mockedService.EXPECT().GetFiles().Return([]string{}, errors.New("an error occurred"))
		req, err := http.NewRequest(http.MethodGet, "/api/v1/files", nil)
		if err != nil {
			t.Fatal(err)
		}
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)
		if !strings.Contains(resp.Body.String(), "an error occurred") {
			t.Fatalf("Expected error message: %s to be in the response body", "an error occurred")
		}
		if resp.Code != http.StatusBadRequest {
			t.Fatalf("Expected status code: %d but got: %d", http.StatusBadRequest, resp.Code)
		}
	})

}
