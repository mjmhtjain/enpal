package dto

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/mjmhtjain/enpal/src/internal/domain"
)

type CalendarQueryRequestBody struct {
	Date     string `json:"date"`
	Products string `json:"products"`
	Language string `json:"language"`
	Rating   string `json:"rating"`
}

func (c *CalendarQueryRequestBody) parseDate() (string, error) {
	t, err := time.Parse(time.RFC3339, c.Date)
	if err != nil {
		return "", errors.New("bad date format")
	}

	return t.Format(time.DateOnly), nil
}

func (c *CalendarQueryRequestBody) parseProducts() ([]domain.Product, error) {
	var productValues []string
	productArr := []domain.Product{}

	// unmarshal the values
	err := json.Unmarshal([]byte(c.Products), &productValues)
	if err != nil {
		return nil, errors.New("bad products")
	}

	// check for empty values
	if len(productValues) == 0 {
		return nil, errors.New("no products")
	}

	// remove duplicates and identify valid values
	validProductMap := domain.GetValidProductsMap()

	for _, p := range productValues {
		val := domain.Product(p)
		if _, exists := validProductMap[val]; exists {
			validProductMap[val] = true
		} else {
			return nil, errors.New("invalid products")
		}
	}

	for key, val := range validProductMap {
		if val {
			productArr = append(productArr, key)
		}
	}

	return productArr, nil
}

func (c *CalendarQueryRequestBody) parseLanguage() (domain.Language, error) {
	var languageVal string

	// unmarshal the values
	err := json.Unmarshal([]byte(c.Language), &languageVal)
	if err != nil {
		return "", errors.New("bad language value")
	}

	return new(domain.Language).GetLanguage(languageVal)
}

func (c *CalendarQueryRequestBody) parseRating() (domain.Rating, error) {
	var ratingVal string

	// unmarshal the values
	err := json.Unmarshal([]byte(c.Rating), &ratingVal)
	if err != nil {
		return "", errors.New("bad rating value")
	}

	return new(domain.Rating).GetRating(ratingVal)
}

// GetDomain validates all the body fields and generates a domain object, otherwise returns an error
func (c *CalendarQueryRequestBody) GetDomainObject() (domain.CalendarQueryDomain, error) {
	date, err := c.parseDate()
	if err != nil {
		return domain.CalendarQueryDomain{}, err
	}

	lang, err := c.parseLanguage()
	if err != nil {
		return domain.CalendarQueryDomain{}, err
	}

	products, err := c.parseProducts()
	if err != nil {
		return domain.CalendarQueryDomain{}, err
	}

	rating, err := c.parseRating()
	if err != nil {
		return domain.CalendarQueryDomain{}, err
	}

	queryObj := domain.CalendarQueryDomain{
		Date:     date,
		Products: products,
		Language: lang,
		Rating:   rating,
	}

	return queryObj, nil
}
