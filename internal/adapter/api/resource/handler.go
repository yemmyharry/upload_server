package resource

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
	domain "upload_server/internal/core/domain/resource"
	"upload_server/internal/core/logger"
)

type HealthCheckResponse struct {
	Status string
}

func (s *HTTPHandler) HealthCheck(c *gin.Context) {
	logger.Info("HealthCheck called")
	c.JSON(200, HealthCheckResponse{Status: "OK"})
}

func (s *HTTPHandler) UploadFile(c *gin.Context) {
	logger.Info("UploadFile called")
	file, err := c.FormFile("file")
	if err != nil {
		logger.Error("No file provided")
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file provided"})
		return
	}

	filename := fmt.Sprintf("%s%s", strconv.FormatInt(file.Size, 10), filepath.Ext(file.Filename))
	if err := c.SaveUploadedFile(file, filename); err != nil {
		logger.Error("Failed to save file")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	_ = s.uploadService.UploadFile(domain.File{Name: filename})

	logger.Info("File uploaded successfully")
	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}
