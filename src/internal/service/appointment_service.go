package service

import (
	"slices"
	"time"

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

func (s *AppointmentService) FindFreeSlots(calQuery domain.CalendarQueryDomain) ([]dto.CalendarQueryResponse, error) {
	response := []dto.CalendarQueryResponse{}
	slots, err := s.appointmentRepo.FindFreeSlots(calQuery.Date)
	if err != nil {
		return nil, err
	}

	// filter the slots based on sales_manager language and rating
	grp := map[time.Time]int{} // time:count

	for _, s := range slots {
		langArr := []string(s.SalesManager.Languages)
		ratingArr := []string(s.SalesManager.CustomerRatings)

		if !slices.Contains(langArr, calQuery.Language.ToString()) {
			continue
		}

		if !slices.Contains(ratingArr, calQuery.Rating.ToString()) {
			continue
		}

		// group these slots based on starttime
		if count, exist := grp[s.StartDate]; exist {
			grp[s.StartDate] = count + 1
		} else {
			grp[s.StartDate] = 1
		}
	}

	// generate response
	for k, v := range grp {
		response = append(response, dto.CalendarQueryResponse{
			AvailableCount: v,
			StartDate:      k.Format(time.RFC3339),
		})
	}

	return response, nil
}
