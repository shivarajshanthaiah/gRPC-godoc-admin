package handler

import (
	"context"
	"fmt"
	"log"

	docpb "github.com/shivaraj-shanthaiah/gRPC/godoc/Admin-Service/pkg/doctor/pb"
	adminpb "github.com/shivaraj-shanthaiah/gRPC/godoc/Admin-Service/pkg/pb"
)

func CreateDoctorHandler(client docpb.DoctorServicesClient, p *adminpb.DoctorModel) (*docpb.DoctorResponse, error) {
	ctx := context.Background()
	fmt.Println(ctx)
	fmt.Println("Hi")
	response, err := client.CreateDoctor(ctx, &docpb.DoctorModel{
		DoctorName: p.DoctorName,
		Age:        p.Age,
		Speciality: p.Speciality,
		Gender:     p.Gender,
		Email:      p.Email,
	})
	fmt.Println(err)
	if err != nil {
		log.Printf("error while creating doctor")
		return nil, err
	}
	return response, nil
}

func FetchDoctorByIDHandler(client docpb.DoctorServicesClient, p *adminpb.DoctorID) (*docpb.DoctorModel, error) {
	ctx := context.Background()

	response, err := client.FetchDoctorByID(ctx, &docpb.DoctorID{Id: p.Id})
	if err != nil {
		log.Printf("error while fetching doctor by name")
		return nil, err
	}
	return response, nil
}

func FetchDoctorByNameHandler(client docpb.DoctorServicesClient, p *adminpb.DoctorName) (*docpb.DoctorModel, error) {
	ctx := context.Background()

	response, err := client.FetchDoctorByName(ctx, &docpb.DoctorName{Name: p.Name})
	if err != nil {
		log.Printf("error while fetching doctor by id")
		return nil, err
	}
	return response, nil
}

func FetchAllDoctorsHandler(client docpb.DoctorServicesClient, p *adminpb.AdminNoParam) (*docpb.DoctorList, error) {
	ctx := context.Background()

	response, err := client.FetchAllDoctors(ctx, &docpb.NoParam{})
	if err != nil {
		log.Printf("error while fetching doctors")
		return nil, err
	}
	return response, nil
}
