package interfaces

import (
	pb "github.com/shivaraj-shanthaiah/gRPC/godoc/Admin-Service/pkg/pb"
)

type AdminServiceInter interface {
	AdminLoginService(p *pb.AdminRequest) (*pb.AdminResponse, error)
	EditUserService(p *pb.UserModel) (*pb.AdminResponse, error)
	DeleteUserService(p *pb.AUserID) (*pb.AdminResponse, error)
	FindUserByEmailService(p *pb.UserRequest) (*pb.User, error)
	FindAllUserService(p *pb.AdminNoParam) (*pb.UserList, error)
	FindUserByIDService(p *pb.AUserID) (*pb.User, error)

	CreateDoctorService(p *pb.DoctorModel) (*pb.AdminResponse, error)
	FetchDoctorByIDService(p *pb.DoctorID) (*pb.DoctorModel, error)
	FetchDoctorByNameService(p *pb.DoctorName) (*pb.DoctorModel, error)
	FetchAllDcotorsService(p *pb.AdminNoParam) (*pb.DoctorList, error)
}
