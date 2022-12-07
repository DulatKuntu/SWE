package service

import (
	"swe/model"
	"swe/pkg/repository"
	"swe/utils"
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
	token, err := utils.GenerateToken(user.ID, utils.Users)
	if err != nil {
		return err
	}
	user.Token = token
	err = s.repo.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
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
	token, err := utils.GenerateToken(doctor.ID, utils.Doctors)
	if err != nil {
		return err
	}
	doctor.Token = token
	return s.repo.CreateDoctor(doctor)
}

func (s *AdminService) GetAllDoctors(search string, specializationID int) ([]*model.DoctorResponse, error) {
	search = "%" + search + "%"
	return s.repo.GetAllDoctors(search, specializationID)
}

func (s *AdminService) GetDoctorById(id string) (*model.Doctor, error) {
	return s.repo.GetDoctorById(id)
}

func (s *AdminService) UpdateDoctor(doctor *model.Doctor) error {
	return s.repo.UpdateDoctor(doctor)
}

func (s *AdminService) CreateSpecialization(specialization *model.Specialization) error {
	return s.repo.CreateSpecialization(specialization)
}

func (s *AdminService) GetAllSpecializations() ([]*model.Specialization, error) {
	return s.repo.GetAllSpecializations()
}

func (s *AdminService) GetSpecializationById(id string) (*model.Specialization, error) {
	return s.repo.GetSpecializationById(id)
}

func (s *AdminService) UpdateSpecialization(specialization *model.Specialization) error {
	return s.repo.UpdateSpecialization(specialization)
}

func (s *AdminService) LoginUser(email, password string) (*model.User, error) {
	return s.repo.LoginUser(email, password)
}

func (s *AdminService) LoginDoctor(iin, password string) (*model.Doctor, error) {
	return s.repo.LoginDoctor(iin, password)
}
