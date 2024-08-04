package handler

import (
	"context"
	"log"

	pb "github.com/shivaraj-shanthaiah/gRPC/godoc/Admin-Service/pkg/pb"
	"github.com/shivaraj-shanthaiah/gRPC/godoc/Admin-Service/pkg/service/interfaces"
)

type AdminHandler struct {
	AdminService interfaces.AdminServiceInter
	pb.AdminServiceServer
}

func NewAdminHandler(repo interfaces.AdminServiceInter) *AdminHandler {
	return &AdminHandler{
		AdminService: repo,
	}
}

func (a *AdminHandler) AdminLogin(ctx context.Context, p *pb.AdminRequest) (*pb.AdminResponse, error) {
	result, err := a.AdminService.AdminLoginService(p)
	if err != nil {
		log.Println("error while logging in")
		return nil, err
	}
	return result, nil
}

func (a *AdminHandler) EditUserFn(ctx context.Context, p *pb.UserModel) (*pb.AdminResponse, error) {
	result, err := a.AdminService.EditUserService(p)
	if err != nil {
		log.Println("error while editing user")
		return nil, err
	}
	return result, nil
}

func (a *AdminHandler) DeleteUserFn(ctx context.Context, p *pb.AUserID) (*pb.AdminResponse, error) {
	result, err := a.AdminService.DeleteUserService(p)
	if err != nil {
		log.Println("error while deleting user")
		return nil, err
	}
	return result, nil
}

func (a *AdminHandler) FindAllUsersFn(ctx context.Context, p *pb.AdminNoParam) (*pb.UserList, error) {
	result, err := a.AdminService.FindAllUserService(p)
	if err != nil {
		log.Println("error while finding all users")
		return nil, err
	}
	return result, nil
}

func (a *AdminHandler) FindUserByEmailFn(ctx context.Context, p *pb.UserRequest) (*pb.User, error) {
	result, err := a.AdminService.FindUserByEmailService(p)
	if err != nil {
		log.Println("error while fetching user")
		return nil, err
	}
	return result, nil
}

func (a *AdminHandler) FindUserByIDFn(ctx context.Context, p *pb.AUserID) (*pb.User, error) {
	result, err := a.AdminService.FindUserByIDService(p)
	if err != nil {
		log.Println("error while fetching user")
		return nil, err
	}
	return result, nil
}
