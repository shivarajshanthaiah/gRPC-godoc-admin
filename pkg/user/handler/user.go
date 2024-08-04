package handler

import (
	"context"
	"log"

	adminpb "github.com/shivaraj-shanthaiah/gRPC/godoc/Admin-Service/pkg/pb"
	userpb "github.com/shivaraj-shanthaiah/gRPC/godoc/Admin-Service/pkg/user/pb"
)

func FindAllUserHandler(client userpb.UserServiceClient, p *adminpb.AdminNoParam) (*userpb.Users, error) {
	ctx := context.Background()
	response, err := client.FindAllUsers(ctx, &userpb.UNoParam{})
	if err != nil {
		log.Printf("error getting all users")
		return nil, err
	}
	return response, nil
}

func EditUserHandler(client userpb.UserServiceClient, p *adminpb.UserModel) (*userpb.CommonResponse, error) {
	ctx := context.Background()

	response, err := client.EditUser(ctx, &userpb.SignupRequest{
		Firstname: p.Firstname,
		Lastname:  p.Lastname,
		Dob:       p.Dob,
		Gender:    p.Gender,
		Email:     p.Email,
		Phone:     p.Phone,
		Password:  p.Password,
	})
	if err != nil {
		log.Printf("error editing user details")
		return nil, err
	}
	return response, nil
}

func DeleteUserHandler(client userpb.UserServiceClient, p *adminpb.AUserID) (*userpb.CommonResponse, error) {
	ctx := context.Background()

	response, err := client.DeleteUser(ctx, &userpb.UserID{
		Id: p.Id,
	})
	if err != nil {

		log.Printf("error while deleting user")
		return nil, err
	}
	return response, nil
}

func FindUserByIDHandler(client userpb.UserServiceClient, p *adminpb.AUserID) (*userpb.Profile, error) {
	ctx := context.Background()

	response, err := client.FindUserByID(ctx, &userpb.UserID{
		Id: p.Id,
	})
	if err != nil {
		log.Printf("error while finding user")
		return nil, err
	}
	return response, nil
}

func FindUserByEmailHandler(client userpb.UserServiceClient, p *adminpb.UserRequest) (*userpb.Profile, error) {
	ctx := context.Background()

	response, err := client.FindUserByEmail(ctx, &userpb.Email{
		Email: p.Email,
	})
	if err != nil {
		log.Printf("error while finding user")
		return nil, err
	}
	return response, nil
}