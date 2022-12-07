package model

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID               string    `json:"id" gorm:"primary_key"`
	Token            string    `json:"token"`
	IIN              string    `json:"iin"`
	Name             string    `json:"name"`
	Surname          string    `json:"surname"`
	MiddleName       string    `json:"middle_name"`
	BloodGroup       string    `json:"blood_group"`
	ContactNumber    string    `json:"contact_number"`
	Email            string    `json:"email"`
	Password         string    `json:"password"`
	Address          string    `json:"address"`
	MaritalStatus    string    `json:"marital_status"`
	RegistrationDate time.Time `json:"registration_date"`
	Records          []*Record `json:"records" gorm:"foreignKey:UserID"`
}

func (p *User) ParseRequest(c *gin.Context) error {
	if err := c.BindJSON(&p); err != nil {
		return errors.New("bad request | " + err.Error())
	}

	return nil
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (p *UserLogin) ParseRequest(c *gin.Context) error {
	if err := c.BindJSON(&p); err != nil {
		return errors.New("bad request | " + err.Error())
	}

	return nil
}

type DoctorLogin struct {
	IIN      string `json:"iin"`
	Password string `json:"password"`
}

func (p *DoctorLogin) ParseRequest(c *gin.Context) error {
	if err := c.BindJSON(&p); err != nil {
		return errors.New("bad request | " + err.Error())
	}

	return nil
}
