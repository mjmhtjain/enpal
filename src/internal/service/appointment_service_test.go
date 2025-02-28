package service

import (
	"testing"
	"time"

	"github.com/mjmhtjain/enpal/src/internal/domain"
	"github.com/mjmhtjain/enpal/src/internal/mocks"
	"github.com/mjmhtjain/enpal/src/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAppointmentService_FindFreeSlots_Success(t *testing.T) {
	// Setup
	mockRepo := mocks.NewAppointmentRepo(t)
	service := NewAppointmentService(mockRepo)

	// Test data
	now := time.Now().Truncate(time.Hour)
	later := now.Add(1 * time.Hour)

	// Create test slots
	slots := []model.Slot{
		{
			ID:        1,
			StartDate: now,
			EndDate:   now.Add(30 * time.Minute),
			SalesManager: model.SalesManager{
				ID:              1,
				Languages:       `["English", "German"]`,
				CustomerRatings: `["Gold", "Bronze"]`,
			},
		},
		{
			ID:        2,
			StartDate: now,
			EndDate:   now.Add(30 * time.Minute),
			SalesManager: model.SalesManager{
				ID:              2,
				Languages:       `["English"]`,
				CustomerRatings: `["Gold"]`,
			},
		},
		{
			ID:        3,
			StartDate: later,
			EndDate:   later.Add(30 * time.Minute),
			SalesManager: model.SalesManager{
				ID:              3,
				Languages:       `["German"]`,
				CustomerRatings: `["Gold"]`,
			},
		},
		{
			ID:        4,
			StartDate: later,
			EndDate:   later.Add(30 * time.Minute),
			SalesManager: model.SalesManager{
				ID:              4,
				Languages:       `["English"]`,
				CustomerRatings: `["Bronze"]`, // Different rating
			},
		},
	}

	// Set up mock expectations
	mockRepo.On("FindFreeSlots", mock.Anything).Return(slots, nil)

	// Query parameters
	query := domain.CalendarQueryDomain{
		Date: time.Now().Format(time.DateOnly),
		Products: []domain.Product{
			domain.Product("SolarPanels"),
			domain.Product("Heatpumps"),
		},
		Language: domain.Language("English"),
		Rating:   domain.Rating("Bronze"),
	}

	// Execute
	results, err := service.FindFreeSlots(query)

	// Verify
	assert.NoError(t, err)
	assert.Len(t, results, 2) // Two distinct time slots

	// Verify the counts are correct
	timeMap := make(map[string]int)
	for _, result := range results {
		timeMap[result.StartDate] = result.AvailableCount
	}

	// 1 slots at the current time match criteria (English + Bronze)
	assert.Equal(t, 1, timeMap[now.Format(time.RFC3339)])

	// 1 slots at the later time match criteria (English + Bronze)
	assert.Equal(t, 1, timeMap[later.Format(time.RFC3339)])

	// Verify mock expectations were met
	mockRepo.AssertExpectations(t)
}

func TestAppointmentService_FindFreeSlots_RepositoryError(t *testing.T) {
	// Setup
	mockRepo := mocks.NewAppointmentRepo(t)
	service := NewAppointmentService(mockRepo)

	// Set up mock to return an error
	expectedError := assert.AnError
	mockRepo.On("FindFreeSlots", mock.Anything).Return(nil, expectedError)

	// Query parameters
	query := domain.CalendarQueryDomain{
		Date: time.Now().Format(time.DateOnly),
		Products: []domain.Product{
			domain.Product("SolarPanels"),
		},
		Language: domain.Language("English"),
		Rating:   domain.Rating("Gold"),
	}

	// Execute
	results, err := service.FindFreeSlots(query)

	// Verify
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Nil(t, results)
	mockRepo.AssertExpectations(t)
}

func TestAppointmentService_FindFreeSlots_NoSlots(t *testing.T) {
	// Setup
	mockRepo := mocks.NewAppointmentRepo(t)
	service := NewAppointmentService(mockRepo)

	// Set up mock to return empty slots
	mockRepo.On("FindFreeSlots", mock.Anything).Return([]model.Slot{}, nil)

	// Query parameters
	query := domain.CalendarQueryDomain{
		Date: time.Now().Format(time.DateOnly),
		Products: []domain.Product{
			domain.Product("SolarPanels"),
		},
		Language: domain.Language("English"),
		Rating:   domain.Rating("Gold"),
	}

	// Execute
	results, err := service.FindFreeSlots(query)

	// Verify
	assert.NoError(t, err)
	assert.Empty(t, results, "Should return empty results when no slots are available")
	mockRepo.AssertExpectations(t)
}

func TestAppointmentService_FindFreeSlots_NoMatchingCriteria(t *testing.T) {
	// Setup
	mockRepo := mocks.NewAppointmentRepo(t)
	service := NewAppointmentService(mockRepo)

	// Test data
	now := time.Now().Truncate(time.Hour)
	slots := []model.Slot{
		{
			ID:        1,
			StartDate: now,
			EndDate:   now.Add(30 * time.Minute),
			SalesManager: model.SalesManager{
				ID:              1,
				Languages:       `["German", "Spanish"]`, // No English language support
				CustomerRatings: `["Gold"]`,
			},
		},
		{
			ID:        2,
			StartDate: now.Add(1 * time.Hour),
			EndDate:   now.Add(90 * time.Minute),
			SalesManager: model.SalesManager{
				ID:              2,
				Languages:       `["English"]`,
				CustomerRatings: `["Silver"]`, // No Gold rating
			},
		},
	}

	// Set up mock expectations
	mockRepo.On("FindFreeSlots", mock.Anything).Return(slots, nil)

	// Query for criteria that won't match any available slots
	query := domain.CalendarQueryDomain{
		Date: time.Now().Format(time.DateOnly),
		Products: []domain.Product{
			domain.Product("SolarPanels"),
		},
		Language: domain.Language("English"),
		Rating:   domain.Rating("Gold"),
	}

	// Execute
	results, err := service.FindFreeSlots(query)

	// Verify
	assert.NoError(t, err)
	assert.Empty(t, results, "Should return empty results when no slots match criteria")
	mockRepo.AssertExpectations(t)
}
