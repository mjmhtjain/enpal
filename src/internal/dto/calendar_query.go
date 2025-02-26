package dto

import "github.com/mjmhtjain/enpal/src/internal/domain"

type CalendarQueryRequestBody struct {
	Date     string `json:"date"`
	Products string `json:"products"`
	Language string `json:"language"`
	Rating   string `json:"rating"`
}

func (c *CalendarQueryRequestBody) parseDate() (string, error) {
	return "", nil
}

func (c *CalendarQueryRequestBody) parseProducts() ([]string, error) {
	return nil, nil
}

func (c *CalendarQueryRequestBody) parseLanguage() (string, error) {
	return "", nil
}

func (c *CalendarQueryRequestBody) parseRating() (string, error) {
	return "", nil
}

// GetDomain validates all the body fields and generates a domain object, otherwise returns an error
func (c *CalendarQueryRequestBody) GetDomain() (domain.CalendarQueryDomain, error) {
	return domain.CalendarQueryDomain{}, nil
}
