package service

import (
	"swe/model"
	"swe/pkg/repository"
	"time"
)

type DoctorService struct {
	repo repository.Doctor
}

func NewDoctorService(repo repository.Doctor) *DoctorService {
	return &DoctorService{repo: repo}
}

func (s *DoctorService) GetAvailableRecords(doctorID string, timeStamp time.Time) ([]time.Time, error) {
	return s.repo.GetAvailableRecords(doctorID, timeStamp)
}

func (s *DoctorService) CreateRecord(record *model.Record) error {
	return s.repo.CreateRecord(record)
}
