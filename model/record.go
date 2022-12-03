package model

import (
	"errors"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type Record struct {
	ID       int       `json:"id" gorm:"primaryKey"`
	DoctorID string    `json:"doctor_id"`
	UserID   string    `json:"user_id"`
	Time     time.Time `json:"time"`
}

func (p *Record) ParseRequest(c *gin.Context) error {
	if err := c.BindJSON(&p); err != nil {
		return errors.New("bad request | " + err.Error())
	}

	return nil
}

type DoctorRecordReq struct {
	DoctorID string    `json:"doctor_id"`
	Time     time.Time `json:"time"`
}

func (p *DoctorRecordReq) ParseRequest(c *gin.Context) error {
	if err := c.BindJSON(&p); err != nil {
		return errors.New("bad request | " + err.Error())
	}
	log.Print(p)

	return nil
}
