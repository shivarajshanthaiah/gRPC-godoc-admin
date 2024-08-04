package service

import (
	"fmt"

	doctor "github.com/shivaraj-shanthaiah/gRPC/godoc/Admin-Service/pkg/doctor/handler"
	adminpb "github.com/shivaraj-shanthaiah/gRPC/godoc/Admin-Service/pkg/pb"
)

func (a *AdminService) CreateDoctorService(p *adminpb.DoctorModel) (*adminpb.AdminResponse, error) {
	fmt.Println("Im here")
	result, err := doctor.CreateDoctorHandler(a.doctorClient, p)
	if err != nil {
		return nil, err
	}

	return &adminpb.AdminResponse{
		Status:  result.Status,
		Error:   result.Error,
		Message: result.Message,
	}, nil
}

func (a *AdminService) FetchDoctorByIDService(p *adminpb.DoctorID) (*adminpb.DoctorModel, error) {
	result, err := doctor.FetchDoctorByIDHandler(a.doctorClient, p)
	if err != nil {
		return nil, err
	}

	return &adminpb.DoctorModel{
		DoctorName: result.DoctorName,
		Age:        result.Age,
		Speciality: result.Speciality,
		Gender:     result.Gender,
		Email:      result.Email,
	}, nil
}

func (a *AdminService) FetchDoctorByNameService(p *adminpb.DoctorName) (*adminpb.DoctorModel, error) {
	result, err := doctor.FetchDoctorByNameHandler(a.doctorClient, p)
	if err != nil {
		return nil, err
	}

	return &adminpb.DoctorModel{
		DoctorName: result.DoctorName,
		Age:        result.Age,
		Speciality: result.Speciality,
		Gender:     result.Gender,
		Email:      result.Email,
	}, nil
}

func (a *AdminService) FetchAllDcotorsService(p *adminpb.AdminNoParam) (*adminpb.DoctorList, error) {
	result, err := doctor.FetchAllDoctorsHandler(a.doctorClient, p)
	if err != nil {
		return nil, err
	}
	var doctors []*adminpb.DoctorModel
	for _, i := range result.Doctors {
		doctors = append(doctors, &adminpb.DoctorModel{
			DoctorName: i.DoctorName,
			Age:        i.Age,
			Speciality: i.Speciality,
			Gender:     i.Gender,
			Email:      i.Email,
		})
	}

	return &adminpb.DoctorList{
		Doctors: doctors,
	}, nil
}
