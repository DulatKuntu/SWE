package model

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

type DoctorResponse struct {
	ID              string    `json:"id" gorm:"primary_key"`
	IIN             string    `json:"iin"`
	DateOfBirth     time.Time `json:"date_of_birth"`
	Name            string    `json:"name"`
	Surname         string    `json:"surname"`
	MiddleName      string    `json:"middle_name"`
	BloodGroup      string    `json:"blood_group"`
	ContactNumber   string    `json:"contact_number"`
	DepartmentID    int       `json:"department_id"`
	Specialization  string    `json:"specialization"`
	Experience      int       `json:"experience"`
	Photo           string    `json:"photo"`
	Category        string    `json:"category"`
	Price           int       `json:"price"`
	ScheduleDetails string    `json:"schedule_details"`
	Degree          string    `json:"degree"`
	Rating          float32   `json:"rating"`
	Address         string    `json:"address"`
	Records         []*Record `json:"records" gorm:"foreignKey:DoctorID"`
}

func (d *DoctorResponse) ReadDoctor(p *Doctor) {
	d.ID = p.ID
	d.IIN = p.IIN
	d.DateOfBirth = p.DateOfBirth
	d.Name = p.Name
	d.Surname = p.Surname
	d.MiddleName = p.MiddleName
	d.BloodGroup = p.BloodGroup
	d.ContactNumber = p.ContactNumber
	d.DepartmentID = p.DepartmentID
	d.Experience = p.Experience
	d.Photo = p.Photo
	d.Category = p.Category
	d.Price = p.Price
	d.ScheduleDetails = p.ScheduleDetails
	d.Degree = p.Degree
	d.Rating = p.Rating
	d.Address = p.Address
	d.Records = p.Records
}

type Doctor struct {
	ID               string    `json:"id" gorm:"primary_key"`
	IIN              string    `json:"iin"`
	Password         string    `json:"password"`
	Token            string    `json:"token"`
	DateOfBirth      time.Time `json:"date_of_birth"`
	Name             string    `json:"name"`
	Surname          string    `json:"surname"`
	MiddleName       string    `json:"middle_name"`
	BloodGroup       string    `json:"blood_group"`
	ContactNumber    string    `json:"contact_number"`
	DepartmentID     int       `json:"department_id"`
	SpecializationID int       `json:"specialization_id" gorm:"foreignKey:SpecializationID"`
	Experience       int       `json:"experience"`
	Photo            string    `json:"photo"`
	Category         string    `json:"category"`
	Price            int       `json:"price"`
	ScheduleDetails  string    `json:"schedule_details"`
	Degree           string    `json:"degree"`
	Rating           float32   `json:"rating"`
	Address          string    `json:"address"`
	Records          []*Record `json:"records" gorm:"foreignKey:DoctorID"`
}

func (p *Doctor) ParseRequest(c *gin.Context) error {
	if err := c.BindJSON(&p); err != nil {
		return errors.New("bad request | " + err.Error())
	}

	if len(p.IIN) != 12 || len(p.ID) != 9 {
		return errors.New("bad request | id is required")
	}

	return nil
}

type SpecializationDoctor struct {
	ID               int    `json:"id" gorm:"primary_key"`
	SpecializationID int    `json:"specialization_id" gorm:"foreignKey:ID"`
	DoctorID         string `json:"doctor_id" gorm:"foreignKey:ID"`
}

type Specialization struct {
	ID      int                     `json:"id" gorm:"primary_key"`
	Name    string                  `json:"name"`
	Doctors []*SpecializationDoctor `json:"doctors" gorm:"ForeignKey:SpecializationID"`
}

func (p *Specialization) ParseRequest(c *gin.Context) error {
	if err := c.BindJSON(&p); err != nil {
		return errors.New("bad request | " + err.Error())
	}

	return nil
}

type DoctorAppointment struct {
	User
	Time time.Time `json:"time"`
}

type UserAppointment struct {
	Doctor
	Time time.Time `json:"time"`
}
