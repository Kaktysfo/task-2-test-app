package events

import (
	"errors"
	"time"

	"github.com/Kaktysfo/app/validation"
	"github.com/araddon/dateparse"
)

type Event struct {
	Title   string
	StartAt time.Time
}

func NewEvent(title string, dateStr string) (Event, error) {
	isValid := validation.IsValidateTitle(title)
	if isValid {
		dateParser, err := dateparse.ParseAny(dateStr)
		if err != nil {
			return Event{}, errors.New("неверный формат даты")
		}
		return Event{
			Title:   title,
			StartAt: dateParser,
		}, nil
	}
	return Event{}, errors.New("неправильный формат имени")
}
