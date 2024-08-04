package user

import (
	"log"

	userpb "github.com/shivaraj-shanthaiah/gRPC/godoc/Admin-Service/pkg/user/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func UserClienDial() (userpb.UserServiceClient, error) {
	grpc, err := grpc.Dial(":8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Error dialing to gRPC client :8082")
		return nil, err
	}

	log.Printf("Successfully Conneced to to client :8082")
	return userpb.NewUserServiceClient(grpc), nil
}
