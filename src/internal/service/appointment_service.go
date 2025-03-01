package service

import (
	"slices"
	"time"

	"github.com/mjmhtjain/enpal/src/internal/domain"
	"github.com/mjmhtjain/enpal/src/internal/dto"
	"github.com/mjmhtjain/enpal/src/internal/repository"
	"github.com/mjmhtjain/enpal/src/internal/util"
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

outerloop:
	for _, s := range slots {
		langArr := []string(s.SalesManager.Languages)
		ratingArr := []string(s.SalesManager.CustomerRatings)
		productArr := []string(s.SalesManager.Products)

		// check for language
		if !slices.Contains(langArr, calQuery.Language.ToString()) {
			continue
		}

		// check for ratings
		if !slices.Contains(ratingArr, calQuery.Rating.ToString()) {
			continue
		}

		// check for products
		productMap := map[string]bool{}
		for _, p := range productArr {
			productMap[p] = true
		}

		for _, p := range calQuery.Products {
			if _, ex := productMap[p.ToString()]; !ex {
				continue outerloop
			}
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
			StartDate:      util.UniversalTimeFormat(k),
		})
	}

	return response, nil
}
