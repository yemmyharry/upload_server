package ports

type UploadService interface {
	HealthCheck() error
	GetFiles() ([]string, error)
}
