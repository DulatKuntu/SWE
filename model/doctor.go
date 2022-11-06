package model

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

type Doctor struct {
	DateOfBirth      time.Time `json:"date_of_birth"`
	IIN              string    `json:"iin"`
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	Surname          string    `json:"surname"`
	MiddleName       string    `json:"middle_name"`
	BloodGroup       string    `json:"blood_group"`
	ContactNumber    string    `json:"contact_number"`
	DepartmentID     int       `json:"department_id"`
	SpecializationID int       `json:"specialization_id"`
	Experience       int       `json:"experience"`
	Photo            string    `json:"photo"`
	Category         string    `json:"category"`
	Price            int       `json:"price"`
	ScheduleDetails  string    `json:"schedule_details"`
	Degree           string    `json:"degree"`
	Rating           float32   `json:"rating"`
	Address          string    `json:"address"`
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
