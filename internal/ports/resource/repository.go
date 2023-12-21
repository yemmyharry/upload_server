package ports

type UploadRepository interface {
	HealthCheck() error
}
