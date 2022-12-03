package repository

import (
	"database/sql"
	"errors"
	"log"
	"swe/model"
	"time"

	"gorm.io/gorm"
)

type DoctorDB struct {
	db     *sql.DB
	gormDB *gorm.DB
}

func NewDoctorDB(db *sql.DB, gormDB *gorm.DB) *DoctorDB {
	return &DoctorDB{db: db, gormDB: gormDB}
}

func (r *DoctorDB) GetAvailableRecords(doctorID string, timeStamp time.Time) ([]time.Time, error) {
	var records []*model.Record
	start := time.Date(timeStamp.Year(), timeStamp.Month(), timeStamp.Day(), 0, 0, 0, 0, timeStamp.Location())
	end := time.Date(timeStamp.Year(), timeStamp.Month(), timeStamp.Day()+1, 0, 0, 0, 0, timeStamp.Location())
	if err := r.gormDB.Table("records").Where("doctor_id = ? AND time >= ?::date AND time <= ?::date", doctorID, start, end).Find(&records).Error; err != nil {
		log.Print(err)
		return nil, err
	}
	var availableRecords []time.Time
	for start.Before(end) {
		var isAvailable = true
		for _, record := range records {
			if record.Time.Year() == start.Year() && record.Time.Month() == start.Month() && record.Time.Day() == start.Day() && record.Time.Hour() == start.Hour() && record.Time.Minute() == start.Minute() {
				isAvailable = false
				break
			}
		}
		if isAvailable {
			availableRecords = append(availableRecords, start)
		}
		start = start.Add(time.Minute * 20)
	}
	return availableRecords, nil
}

func (r *DoctorDB) CreateRecord(record *model.Record) error {
	oldRecord := &model.Record{}
	if err := r.gormDB.Model(&model.Record{}).Where("doctor_id = ? AND time = ?", record.DoctorID, record.Time).First(&oldRecord).Error; err != nil {
		if r.gormDB.Create(record).Error.Error() != "record not found" {
			return err
		}
	}
	log.Print(oldRecord)
	if oldRecord != nil {
		return errors.New("already exists")
	}
	return r.gormDB.Create(record).Error
}
