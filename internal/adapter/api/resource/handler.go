package resource

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
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
	filePath := filepath.Join("files", filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		logger.Error("Error saving file: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	logger.Info("File uploaded successfully")
	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}

func (s *HTTPHandler) GetAllFiles(c *gin.Context) {
	logger.Info("Get files called")

	files, err := s.uploadService.GetFiles()
	if err != nil {
		logger.Error("Error getting files: " + err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"List of files": files})
}

func (s *HTTPHandler) DownloadFile(c *gin.Context) {
	logger.Info("Download file called")
	filename := c.Param("filename")
	filePath := filepath.Join("files", filename)
	c.File(filePath)
}
