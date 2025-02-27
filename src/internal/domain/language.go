package domain

import "errors"

type Language string

const (
	german  = "German"
	english = "English"
)

func (l *Language) GetLanguage(lang string) (Language, error) {
	switch lang {
	case german:
		return german, nil
	case english:
		return english, nil
	default:
		return "", errors.New("invalid language")
	}
}
