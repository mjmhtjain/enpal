package model

type AppointmentGroup struct {
	ID    int `json:"id"`
	Count int `json:"count_slots"`
}
