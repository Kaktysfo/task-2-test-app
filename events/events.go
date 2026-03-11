package events

import (
	"errors"
	"time"

	"github.com/Kaktysfo/app/validation"
	"github.com/araddon/dateparse"
	"github.com/google/uuid"
)

type Event struct {
	ID      string    `json:"id"`
	Title   string    `json:"title"`
	StartAt time.Time `json:"-"`
}

func getNextID() string {
	return uuid.New().String()
}

func (e *Event) Update(title, dateStr string) error {
	date, dateErr := dateparse.ParseAny(dateStr)
	if dateErr != nil {
		return dateErr
	}
	if validation.IsValidateTitle(title) {
		e.Title = title
		e.StartAt = date
	} else {
		return errors.New("неверный заголовок")
	}
	return nil
}

func NewEvent(title string, dateStr string) (*Event, error) {
	isValid := validation.IsValidateTitle(title)
	if isValid {
		dateParser, err := dateparse.ParseAny(dateStr)
		if err != nil {
			return &Event{}, errors.New("неверный формат даты")
		}
		return &Event{
			ID:      getNextID(),
			Title:   title,
			StartAt: dateParser,
		}, nil
	}
	return &Event{}, errors.New("неправильный формат имени")
}
