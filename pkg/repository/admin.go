package repository

import (
	"database/sql"
	"log"
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

func (r *AdminDB) GetAllDoctors(search string, specializationID int) ([]*model.DoctorResponse, error) {
	var doctors []*model.Doctor
	var specialization model.Specialization
	var doctorsResponses []*model.DoctorResponse

	if specializationID == 0 {
		err := r.gormDB.Where("name LIKE ? OR surname LIKE ?", search, search).Find(&doctors).Error
		if err != nil {
			return nil, err
		}
		log.Print(doctors)
		for _, doctor := range doctors {
			doctorResponses := &model.DoctorResponse{}
			doctorResponses.ReadDoctor(doctor)
			log.Print(doctorResponses)
			err = r.gormDB.Where("id = ?", doctor.SpecializationID).Find(&specialization).Error
			if err != nil {
				return nil, err
			}
			doctorResponses.Specialization = specialization.Name
			doctorsResponses = append(doctorsResponses, doctorResponses)
		}
	} else {
		err := r.gormDB.Where("name LIKE ? OR surname LIKE ? AND specialization_id = ?", search, search, specializationID).Find(&doctors).Error
		if err != nil {
			return nil, err
		}
		for _, doctor := range doctors {
			doctorResponses := &model.DoctorResponse{}
			doctorResponses.ReadDoctor(doctor)
			err = r.gormDB.Where("id = ?", doctor.SpecializationID).Find(&specialization).Error
			if err != nil {
				return nil, err
			}
			doctorResponses.Specialization = specialization.Name
			doctorsResponses = append(doctorsResponses, doctorResponses)
		}
	}
	if len(doctorsResponses) == 0 {
		doctorsResponses = []*model.DoctorResponse{}
	}
	return doctorsResponses, r.gormDB.Find(&doctors).Error
}

func (r *AdminDB) GetDoctorById(id string) (*model.Doctor, error) {
	var doctor model.Doctor

	return &doctor, r.gormDB.First(&doctor, id).Error
}

func (r *AdminDB) UpdateDoctor(doctor *model.Doctor) error {
	return r.gormDB.Save(doctor).Error
}

func (r *AdminDB) CreateSpecialization(specialization *model.Specialization) error {
	return r.gormDB.Create(specialization).Error
}

func (r *AdminDB) GetAllSpecializations() ([]*model.Specialization, error) {
	var specializations []*model.Specialization

	return specializations, r.gormDB.Find(&specializations).Error
}

func (r *AdminDB) GetSpecializationById(id string) (*model.Specialization, error) {
	var specialization model.Specialization

	return &specialization, r.gormDB.First(&specialization, id).Error
}

func (r *AdminDB) UpdateSpecialization(specialization *model.Specialization) error {
	return r.gormDB.Save(specialization).Error
}
