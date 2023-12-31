package resource

import (
	"github.com/gin-gonic/gin"
)

func (s *HTTPHandler) Routes(router *gin.Engine) {
	apirouter := router.Group("api/v1")
	apirouter.GET("/healthcheck", s.HealthCheck)
	apirouter.POST("/upload", s.UploadFile)
	apirouter.GET("/files", s.GetAllFiles)
	apirouter.GET("/download/:filename", s.DownloadFile)
	router.NoRoute(func(c *gin.Context) { c.JSON(404, "no route") })
}
