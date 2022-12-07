package service

import (
	"swe/model"
	"swe/pkg/repository"
	"time"
)

type Admin interface {
	CreateUser(user *model.User) error
	GetAllUsers() ([]*model.User, error)
	GetUserById(id string) (*model.User, error)
	UpdateUser(user *model.User) error
	CreateDoctor(doctor *model.Doctor) error
	GetAllDoctors(search string, specializationID int) ([]*model.DoctorResponse, error)
	GetDoctorById(id string) (*model.Doctor, error)
	UpdateDoctor(doctor *model.Doctor) error
	CreateSpecialization(specialization *model.Specialization) error
	GetAllSpecializations() ([]*model.Specialization, error)
	GetSpecializationById(id string) (*model.Specialization, error)
	UpdateSpecialization(specialization *model.Specialization) error
	LoginUser(email, password string) (*model.User, error)
	LoginDoctor(iin, password string) (*model.Doctor, error)
}

type Doctor interface {
	GetAvailableRecords(doctorID string, timeStamp time.Time) ([]time.Time, error)
	CreateRecord(record *model.Record) error
	GetDoctorAppointments(id string) ([]*model.DoctorAppointment, error)
	GetUserAppointments(id string) ([]*model.UserAppointment, error)
}

type Service struct {
	Admin
	Doctor
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Admin:  NewAdminService(repos),
		Doctor: NewDoctorService(repos),
	}
}
