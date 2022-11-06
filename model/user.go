package model

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	IIN              string    `json:"iin"`
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	Surname          string    `json:"surname"`
	MiddleName       string    `json:"middle_name"`
	BloodGroup       string    `json:"blood_group"`
	ContactNumber    string    `json:"contact_number"`
	Email            string    `json:"email"`
	Address          string    `json:"address"`
	MaritalStatus    string    `json:"marital_status"`
	RegistrationDate time.Time `json:"registration_date"`
}

func (p *User) ParseRequest(c *gin.Context) error {
	if err := c.BindJSON(&p); err != nil {
		return errors.New("bad request | " + err.Error())
	}

	if len(p.IIN) != 12 || len(p.ID) != 9 {
		return errors.New("bad request | id is required")
	}

	return nil
}
