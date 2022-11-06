package repository

import (
	"database/sql"
	"swe/model"

	"gorm.io/gorm"
)

type AdminDB struct {
	db     *sql.DB
	gormDB *gorm.DB
}

func NewAdminDB(db *sql.DB, gormDB *gorm.DB) *AdminDB {
	return &AdminDB{db: db, gormDB: gormDB}
}

func (r *AdminDB) CreateUser(user *model.User) error {
	return r.gormDB.Create(user).Error
}

func (r *AdminDB) GetAllUsers() ([]*model.User, error) {
	var users []*model.User

	return users, r.gormDB.Find(&users).Error
}

func (r *AdminDB) GetUserById(id string) (*model.User, error) {
	var user model.User

	return &user, r.gormDB.First(&user, id).Error
}

func (r *AdminDB) UpdateUser(user *model.User) error {
	return r.gormDB.Save(user).Error
}

func (r *AdminDB) CreateDoctor(doctor *model.Doctor) error {
	return r.gormDB.Create(doctor).Error
}

func (r *AdminDB) GetAllDoctors() ([]*model.Doctor, error) {
	var doctors []*model.Doctor

	return doctors, r.gormDB.Find(&doctors).Error
}

func (r *AdminDB) GetDoctorById(id string) (*model.Doctor, error) {
	var doctor model.Doctor

	return &doctor, r.gormDB.First(&doctor, id).Error
}

func (r *AdminDB) UpdateDoctor(doctor *model.Doctor) error {
	return r.gormDB.Save(doctor).Error
}
