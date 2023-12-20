package resource

import (
	"github.com/gin-gonic/gin"
	"upload_server/internal/core/logger"
)

type HealthCheckResponse struct {
	Status string
}

func (s *HTTPHandler) HealthCheck(c *gin.Context) {
	logger.Info("HealthCheck called")
	c.JSON(200, HealthCheckResponse{Status: "OK"})
}
