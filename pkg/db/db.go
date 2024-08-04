package db

import (
	"log"
	"os"

	"github.com/shivaraj-shanthaiah/gRPC/godoc/Admin-Service/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn, ok := os.LookupEnv("DB_Config")
	if !ok {
		log.Fatal("Error Loading Admin database env")
	}

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connnecting to admin database %v", err.Error())
	}

	DB.AutoMigrate(&models.AdminModel{})
	return DB
}