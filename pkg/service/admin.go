package service

import (
	"swe/model"
	"swe/pkg/repository"
	"time"
)

type AdminService struct {
	repo repository.Admin
}

func NewAdminService(repo repository.Admin) *AdminService {
	return &AdminService{repo: repo}
}

func (s *AdminService) CreateUser(user *model.User) error {
	user.RegistrationDate = time.Now()
	return s.repo.CreateUser(user)
}

func (s *AdminService) GetAllUsers() ([]*model.User, error) {
	return s.repo.GetAllUsers()
}

func (s *AdminService) GetUserById(id string) (*model.User, error) {
	return s.repo.GetUserById(id)
}

func (s *AdminService) UpdateUser(user *model.User) error {
	return s.repo.UpdateUser(user)
}

func (s *AdminService) CreateDoctor(doctor *model.Doctor) error {
	return s.repo.CreateDoctor(doctor)
}

func (s *AdminService) GetAllDoctors() ([]*model.Doctor, error) {
	return s.repo.GetAllDoctors()
}

func (s *AdminService) GetDoctorById(id string) (*model.Doctor, error) {
	return s.repo.GetDoctorById(id)
}

func (s *AdminService) UpdateDoctor(doctor *model.Doctor) error {
	return s.repo.UpdateDoctor(doctor)
}
