package model

import "time"

type Record struct {
	ID       int       `json:"id" gorm:"primaryKey"`
	DoctorID int       `json:"doctor_id"`
	UserID   int       `json:"user_id"`
	Time     time.Time `json:"time"`
}
