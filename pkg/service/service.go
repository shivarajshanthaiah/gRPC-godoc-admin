package service

import (
	"errors"

	docpb "github.com/shivaraj-shanthaiah/gRPC/godoc/Admin-Service/pkg/doctor/pb"
	pb "github.com/shivaraj-shanthaiah/gRPC/godoc/Admin-Service/pkg/pb"
	inter "github.com/shivaraj-shanthaiah/gRPC/godoc/Admin-Service/pkg/repo/interface"
	"github.com/shivaraj-shanthaiah/gRPC/godoc/Admin-Service/pkg/service/interfaces"
	user "github.com/shivaraj-shanthaiah/gRPC/godoc/Admin-Service/pkg/user/handler"
	userpb "github.com/shivaraj-shanthaiah/gRPC/godoc/Admin-Service/pkg/user/pb"
)

type AdminService struct {
	AdminRepo    inter.AdminRepoInterface
	client       userpb.UserServiceClient
	doctorClient docpb.DoctorServicesClient
}

func NewAdminService(repo inter.AdminRepoInterface, client userpb.UserServiceClient, doctorClient docpb.DoctorServicesClient) interfaces.AdminServiceInter {
	return &AdminService{
		AdminRepo:    repo,
		client:       client,
		doctorClient: doctorClient,
	}
}

func (a *AdminService) AdminLoginService(p *pb.AdminRequest) (*pb.AdminResponse, error) {
	admin, err := a.AdminRepo.FindAdmin(p.Username)
	if err != nil {
		return nil, err
	}
	if admin.Password != p.Password {
		return nil, errors.New("incorrect password")
	}

	token, err := a.GenerateToken(p.Username)
	if err != nil {
		return nil, err
	}

	response := pb.AdminResponse{
		Status:  "Success",
		Error:   "",
		Message: token,
	}
	return &response, nil
}

func (a *AdminService) EditUserService(p *pb.UserModel) (*pb.AdminResponse, error) {
	result, err := user.EditUserHandler(a.client, p)
	if err != nil {
		return nil, err
	}
	response := &pb.AdminResponse{
		Status:  result.Status,
		Error:   result.Error,
		Message: result.Message,
	}
	return response, nil
}

func (a *AdminService) DeleteUserService(p *pb.AUserID) (*pb.AdminResponse, error) {
	result, err := user.DeleteUserHandler(a.client, p)
	if err != nil {
		return nil, err
	}
	respose := &pb.AdminResponse{
		Status:  result.Status,
		Error:   result.Error,
		Message: result.Message,
	}

	return respose, nil
}

func (a *AdminService) FindUserByEmailService(p *pb.UserRequest) (*pb.User, error) {
	result, err := user.FindUserByEmailHandler(a.client, p)
	if err != nil {
		return nil, err
	}

	return &pb.User{
		Firstname: result.Firstname,
		Lastname:  result.Lastname,
		Username:  result.Username,
		Email:     result.Email,
		Dob:       result.Dob,
		Gender:    result.Gender,
		Phone:     result.Phone,
	}, nil
}

func (a *AdminService) FindAllUserService(p *pb.AdminNoParam) (*pb.UserList, error) {
	result, err := user.FindAllUserHandler(a.client, p)
	if err != nil {
		return nil, err
	}

	var users []*pb.User
	for _, i := range result.Userlist {
		users = append(users, &pb.User{
			Firstname: i.Firstname,
			Lastname:  i.Lastname,
			Username:  i.Username,
			Email:     i.Email,
			Dob:       i.Dob,
			Gender:    i.Gender,
			Phone:     i.Phone,
		})
	}
	response := &pb.UserList{
		Users: users,
	}

	return response, nil
}

func (a *AdminService) FindUserByIDService(p *pb.AUserID) (*pb.User, error) {
	result, err := user.FindUserByIDHandler(a.client, p)
	if err != nil {
		return nil, err
	}

	return &pb.User{
		Firstname: result.Firstname,
		Lastname:  result.Lastname,
		Username:  result.Username,
		Email:     result.Email,
		Dob:       result.Dob,
		Gender:    result.Gender,
		Phone:     result.Phone,
	}, nil
}
