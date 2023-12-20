package resource

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	_ "upload_server/internal/core/domain/resource"
)

type UploadRepositoryDB struct {
	db *gorm.DB
}

var _ = godotenv.Load("file_upload.env")

var DbUsername = os.Getenv("DB_USER")
var DbPassword = os.Getenv("DB_PASS")
var DbName = os.Getenv("DB_NAME")
var DbHost = os.Getenv("DB_HOST")
var DbPort = os.Getenv("DB_PORT")

func NewUploadRepositoryDB() *UploadRepositoryDB {
	dsn := DbUsername + ":" + DbPassword + "@tcp" + "(" + DbHost + ":" + DbPort + ")/" + DbName + "?" + "charset=utf8mb4&parseTime=True&loc=Local"
	client, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	//client.AutoMigrate(&domain.File{})
	return &UploadRepositoryDB{client}
}
