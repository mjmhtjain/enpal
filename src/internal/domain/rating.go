package domain

import "errors"

type Rating string

const (
	gold   = "Gold"
	silver = "Silver"
	bronze = "Bronze"
)

func (l *Rating) GetRating(val string) (Rating, error) {
	switch val {
	case gold:
		return gold, nil
	case silver:
		return silver, nil
	case bronze:
		return bronze, nil
	default:
		return "", errors.New("invalid rating")
	}
}
