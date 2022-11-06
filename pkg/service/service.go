package service

import (
	"swe/model"
	"swe/pkg/repository"
)

type Admin interface {
	CreateUser(user *model.User) error
	GetAllUsers() ([]*model.User, error)
	GetUserById(id string) (*model.User, error)
	UpdateUser(user *model.User) error
	CreateDoctor(doctor *model.Doctor) error
	GetAllDoctors() ([]*model.Doctor, error)
	GetDoctorById(id string) (*model.Doctor, error)
	UpdateDoctor(doctor *model.Doctor) error
}

type Service struct {
	Admin
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Admin: NewAdminService(repos),
	}
}
