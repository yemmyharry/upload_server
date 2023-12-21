package resource

import (
	ports "upload_server/internal/ports/resource"
)

type HTTPHandler struct {
	uploadService ports.UploadService
}

func NewHTTPHandler(uploadService ports.UploadService) *HTTPHandler {
	handler := &HTTPHandler{
		uploadService: uploadService,
	}
	return handler
}
