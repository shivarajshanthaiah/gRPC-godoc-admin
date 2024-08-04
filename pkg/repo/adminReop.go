package repo

import (
	"github.com/shivaraj-shanthaiah/gRPC/godoc/Admin-Service/pkg/models"
	inter "github.com/shivaraj-shanthaiah/gRPC/godoc/Admin-Service/pkg/repo/interface"
	"gorm.io/gorm"
)

type AdminRepo struct {
	DB *gorm.DB
}

func NewAdminRepo(db *gorm.DB) inter.AdminRepoInterface {
	return &AdminRepo{
		DB: db,
	}
}

func (a *AdminRepo) FindAdmin(username string) (*models.AdminModel, error) {
	var admin models.AdminModel
	err := a.DB.Where("username = ?", username).First(&admin).Error
	return &admin, err
}
