package events

import (
	"errors"
	"github.com/araddon/dateparse"
	"regexp"
	"time"
)

type Event struct {
	Title   string
	StartAt time.Time
}

func NewEvent(title string, dateStr string) (Event, error) {
	dateParser, err := dateparse.ParseAny(dateStr)
	if err != nil {
		return Event{}, errors.New("неверный формат даты")
	}
	return Event{
		Title:   title,
		StartAt: dateParser,
	}, nil
}

func isValidateTitle(tilte string) bool {
	pattern := "^[a-zA-Z0-9 ]{3,250}$"
	matched, err := regexp.MatchString(pattern, tilte)
	if err != nil {
		return false
	}
	return matched
}
