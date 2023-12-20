package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	adapter "upload_server/internal/adapter/api/resource"
	"upload_server/internal/adapter/repositories/mysql/resource"
	"upload_server/internal/core/logger"
	services "upload_server/internal/core/services/resource"
)

func main() {
	err := godotenv.Load("file_upload.env")
	if err != nil {
		logger.Error("Error loading .env file")
	}
	router := gin.Default()
	database := resource.NewUploadRepositoryDB()
	service := services.New(database)
	handler := adapter.NewHTTPHandler(service)
	handler.Routes(router)
	logger.Info("Starting server on port 8088")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8088"
	}
	router.Run(":" + port)
}
