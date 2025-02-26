package service

import (
	"github.com/mjmhtjain/enpal/src/internal/domain"
	"github.com/mjmhtjain/enpal/src/internal/dto"
	"github.com/mjmhtjain/enpal/src/internal/repository"
)

type IAppointmentService interface {
	FindFreeSlots(domain.CalendarQueryDomain) ([]dto.CalendarQueryResponse, error)
}

type AppointmentService struct {
	appointmentRepo repository.IAppointmentRepo
}

func NewAppointmentService(appointmentRepo repository.IAppointmentRepo) *AppointmentService {
	return &AppointmentService{
		appointmentRepo: appointmentRepo,
	}
}

func (s *AppointmentService) FindFreeSlots(domain.CalendarQueryDomain) ([]dto.CalendarQueryResponse, error) {
	return nil, nil
}
