package repository

import (
	"database/sql"
	"log"

	"github.com/mjmhtjain/enpal/src/internal/client"
	"github.com/mjmhtjain/enpal/src/internal/model"
)

type IAppointmentRepo interface {
	FindFreeSlots(startDate string) ([]model.AppointmentGroup, error)
}

type AppointmentRepo struct {
	db *sql.DB
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

func (a *AppointmentRepo) FindFreeSlots(startDate string) ([]model.AppointmentGroup, error) {
	result := []model.AppointmentGroup{}

	query := `
		SELECT sa.id, COUNT(*) as count_slots from "slots" sl
	INNER JOIN "sales_managers" sa
	ON sl.sales_manager_id = sa.id
	where sl.booked = 'false' AND
	DATE(sl.start_date) = $1 
	GROUP BY sa.id
	`
	rows, err := a.db.Query(query, startDate)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var appGrp model.AppointmentGroup

		err := rows.Scan(
			&appGrp.ID,
			&appGrp.Count,
		)
		if err != nil {
			return nil, err
		}

		result = append(result, appGrp)
	}

	return result, err
}
