package service

import (
	"log"
	"sort"
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

func (s *DoctorService) GetDoctorAppointments(id string) ([]*model.DoctorAppointment, error) {
	res, err := s.repo.GetDoctorAppointments(id)
	if err != nil {
		return nil, err
	}
	log.Print(res)
	sort.Slice(res, func(i, j int) bool {
		return res[i].Time.Before(res[j].Time)
	})
	return res, nil
}

func (s *DoctorService) GetUserAppointments(id string) ([]*model.UserAppointment, error) {
	res, err := s.repo.GetUserAppointments(id)
	if err != nil {
		return nil, err
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].Time.Before(res[j].Time)
	})
	return res, nil
}
