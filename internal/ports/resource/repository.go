package ports

type UploadRepository interface {
	HealthCheck() error
	GetFiles() ([]string, error)
}
