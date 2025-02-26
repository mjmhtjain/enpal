package dto

type CalendarQueryResponse struct {
	AvailableCount int    `json:"available_count"`
	StartDate      string `json:"start_date"`
}
