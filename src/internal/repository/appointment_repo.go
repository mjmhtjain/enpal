package repository

import (
	"log"

	"github.com/mjmhtjain/enpal/src/internal/client"
	"github.com/mjmhtjain/enpal/src/internal/model"
	"gorm.io/gorm"
)

type IAppointmentRepo interface {
	FindSlots(startDate string) ([]model.Slot, error)
}

type AppointmentRepo struct {
	db *gorm.DB
}

func NewAppointmentRepo() IAppointmentRepo {
	db, err := client.NewDBClient(client.NewDatabaseConfig())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err.Error())
	}

	return &AppointmentRepo{
		db: db,
	}
}

func (a *AppointmentRepo) FindSlots(startDate string) ([]model.Slot, error) {
	var slots []model.Slot

	err := a.db.Preload("SalesManager").
		Where("DATE(start_date) = ?", startDate).
		Find(&slots).Error
	if err != nil {
		return nil, err
	}

	return slots, nil
}
