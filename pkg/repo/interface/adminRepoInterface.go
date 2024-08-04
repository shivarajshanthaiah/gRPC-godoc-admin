package interfaces

import "github.com/shivaraj-shanthaiah/gRPC/godoc/Admin-Service/pkg/models"

type AdminRepoInterface interface {
	FindAdmin(username string) (*models.AdminModel, error)
}
