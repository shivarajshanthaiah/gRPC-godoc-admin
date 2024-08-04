package handler

import (
	"context"
	"log"

	adminpb "github.com/shivaraj-shanthaiah/gRPC/godoc/Admin-Service/pkg/pb"
)

func (a *AdminHandler) CreateDoctor(ctx context.Context, p *adminpb.DoctorModel) (*adminpb.AdminResponse, error) {
	result, err := a.AdminService.CreateDoctorService(p)
	if err != nil {
		log.Println("error while creating doctor")
		return nil, err
	}
	return result, nil
}

func (a *AdminHandler) FetchDoctorByID(ctx context.Context, p *adminpb.DoctorID) (*adminpb.DoctorModel, error) {
	result, err := a.AdminService.FetchDoctorByIDService(p)
	if err != nil {
		log.Println("error while fetching doctor by id")
		return nil, err
	}
	return result, nil
}

func (a *AdminHandler) FetchDoctorByName(ctx context.Context, p *adminpb.DoctorName) (*adminpb.DoctorModel, error) {
	result, err := a.AdminService.FetchDoctorByNameService(p)
	if err != nil {
		log.Println("error while fetching doctor by name")
		return nil, err
	}
	return result, nil
}

func (a *AdminHandler) FetchAllDoctors(ctx context.Context, p *adminpb.AdminNoParam) (*adminpb.DoctorList, error) {
	result, err := a.AdminService.FetchAllDcotorsService(p)
	if err != nil {
		log.Println("error while fetching all doctor")
		return nil, err
	}
	return result, nil
}
