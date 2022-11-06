package repository

import (
	"database/sql"
	"swe/model"

	"gorm.io/gorm"
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

type Repository struct {
	Admin
}

func NewRepository(db *sql.DB, gormDB *gorm.DB) *Repository {
	return &Repository{
		Admin: NewAdminDB(db, gormDB),
	}
}
