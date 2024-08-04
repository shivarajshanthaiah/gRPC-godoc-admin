package di

import (
	"log"

	"github.com/shivaraj-shanthaiah/gRPC/godoc/Admin-Service/config"
	"github.com/shivaraj-shanthaiah/gRPC/godoc/Admin-Service/pkg/db"
	"github.com/shivaraj-shanthaiah/gRPC/godoc/Admin-Service/pkg/doctor"
	"github.com/shivaraj-shanthaiah/gRPC/godoc/Admin-Service/pkg/handler"
	"github.com/shivaraj-shanthaiah/gRPC/godoc/Admin-Service/pkg/repo"
	"github.com/shivaraj-shanthaiah/gRPC/godoc/Admin-Service/pkg/server"
	"github.com/shivaraj-shanthaiah/gRPC/godoc/Admin-Service/pkg/service"
	"github.com/shivaraj-shanthaiah/gRPC/godoc/Admin-Service/pkg/user"
)

func Init() {
	config.LoadConfig()
	db := db.ConnectDB()
	userClient, err := user.UserClienDial()
	if err != nil {
		log.Fatalf("something went wrong when dialing user client %v", err)
	}
	doctorClient, err := doctor.ClientDial()
	if err != nil {
		log.Fatalf("something went wrong when dialing book client %v", err)
	}

	adminRepo := repo.NewAdminRepo(db)
	adminService := service.NewAdminService(adminRepo, userClient, doctorClient)
	adminHandler := handler.NewAdminHandler(adminService)
	server.NewGrpcServer(adminHandler)
}
