package ports

type UploadService interface {
	HealthCheck() error
}
