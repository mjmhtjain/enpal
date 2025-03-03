package service

import (
	"slices"
	"sort"
	"time"

	"github.com/mjmhtjain/enpal/src/internal/domain"
	"github.com/mjmhtjain/enpal/src/internal/dto"
	"github.com/mjmhtjain/enpal/src/internal/model"
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

func (a *AppointmentService) FindFreeSlots(calQuery domain.CalendarQueryDomain) ([]dto.CalendarQueryResponse, error) {
	response := []dto.CalendarQueryResponse{}
	bookedSlots := map[int][]model.Slot{}
	freeSlots := []model.Slot{}
	slots, err := a.appointmentRepo.FindSlots(calQuery.Date)
	if err != nil {
		return nil, err
	}

	// arrange the slots based on starttime
	sort.Slice(slots, func(i, j int) bool {
		return slots[i].StartDate.Before(slots[j].StartDate)
	})

	// filter the slots based on sales_manager language and rating
	grp := map[time.Time]int{} // time:count

	for _, s := range slots {
		// create a map of booked slots for the same sales manager
		if s.Booked {
			if bs, ex := bookedSlots[s.SalesManager.ID]; ex {
				bs = append(bs, s)
				bookedSlots[s.SalesManager.ID] = bs
			} else {
				bookedSlots[s.SalesManager.ID] = []model.Slot{s}
			}
		} else {
			freeSlots = append(freeSlots, s)
		}
	}

outerloop:
	for _, s := range freeSlots {
		langArr := []string(s.SalesManager.Languages)
		ratingArr := []string(s.SalesManager.CustomerRatings)
		productArr := []string(s.SalesManager.Products)

		// check for language
		if !slices.Contains(langArr, calQuery.Language.ToString()) {
			continue outerloop
		}

		// check for ratings
		if !slices.Contains(ratingArr, calQuery.Rating.ToString()) {
			continue outerloop
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

		//check for overlapping slots with booked slots for the same sales manager
		if bs, ex := bookedSlots[s.SalesManager.ID]; ex {
			idx := a.binarySearch(bs, s.StartDate)

			if idx == len(bs) {
				// check if the slot is overlapping with the last booked slot
				if !(s.StartDate.After(bs[idx-1].EndDate) ||
					s.StartDate.Equal(bs[idx-1].EndDate)) {
					continue outerloop
				}
			} else if idx == 0 {
				// check if the slot is overlapping with the first booked slot
				if !(s.EndDate.Before(bs[idx].StartDate) ||
					s.EndDate.Equal(bs[idx].StartDate)) {
					continue outerloop
				}
			} else {
				// check if the slot is overlapping with the previous and next booked slot
				if !(s.EndDate.Before(bs[idx].StartDate) ||
					s.EndDate.Equal(bs[idx].StartDate)) ||
					!(s.StartDate.After(bs[idx-1].EndDate) ||
						s.StartDate.Equal(bs[idx-1].EndDate)) {
					continue outerloop
				}
			}
		}

		// group these slots based on starttime
		if count, exist := grp[s.StartDate]; exist {
			grp[s.StartDate] = count + 1
		} else {
			grp[s.StartDate] = 1
		}
	}

	// sort the grpList based on start date
	type grpType struct {
		StartDate time.Time
		Count     int
	}

	grpList := []grpType{}
	for k, v := range grp {
		grpList = append(grpList, grpType{
			StartDate: k,
			Count:     v,
		})
	}

	sort.Slice(grpList, func(i, j int) bool {
		return grpList[i].StartDate.Before(grpList[j].StartDate)
	})

	// generate response
	for _, v := range grpList {
		response = append(response, dto.CalendarQueryResponse{
			AvailableCount: v.Count,
			StartDate:      util.UniversalTimeFormat(v.StartDate),
		})
	}

	return response, nil
}

// binary search to find the index of the slot with the given start date
func (a *AppointmentService) binarySearch(slots []model.Slot, targetDate time.Time) int {
	low := 0
	high := len(slots) - 1

	for low <= high {
		mid := low + (high-low)/2

		if slots[mid].StartDate.Equal(targetDate) {
			return mid
		} else if slots[mid].StartDate.Before(targetDate) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return low
}
