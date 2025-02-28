package dto

import (
	"errors"
	"time"

	"github.com/mjmhtjain/enpal/src/internal/domain"
)

type CalendarQueryRequestBody struct {
	Date     string   `json:"date"`
	Products []string `json:"products"`
	Language string   `json:"language"`
	Rating   string   `json:"rating"`
}

func (c *CalendarQueryRequestBody) parseDate() (string, error) {
	t, err := time.Parse(time.DateOnly, c.Date)
	if err != nil {
		return "", errors.New("bad date format")
	}

	return t.Format(time.DateOnly), nil
}

func (c *CalendarQueryRequestBody) parseProducts() ([]domain.Product, error) {
	// var productValues []string
	productArr := []domain.Product{}

	// unmarshal the values
	// err := json.Unmarshal([]byte(c.Products), &productValues)
	// if err != nil {
	// 	return nil, errors.New("bad products")
	// }

	// check for empty values
	if len(c.Products) == 0 {
		return nil, errors.New("no products")
	}

	// remove duplicates and identify valid values
	validProductMap := domain.GetValidProductsMap()

	for _, p := range c.Products {
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
	return new(domain.Language).GetLanguage(c.Language)
}

func (c *CalendarQueryRequestBody) parseRating() (domain.Rating, error) {
	return new(domain.Rating).GetRating(c.Rating)
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
