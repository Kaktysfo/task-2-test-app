package events

import (
	"errors"
	"time"

	"github.com/Kaktysfo/app/reminder"
	"github.com/Kaktysfo/app/validation"
	"github.com/araddon/dateparse"
	"github.com/google/uuid"
)

type Event struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	StartAt  time.Time `json:"-"`
	Priority Priority
	Reminder *reminder.Reminder
}

func getNextID() string {
	return uuid.New().String()
}

func (e *Event) Update(title, dateStr, priority string) error {
	date, dateErr := dateparse.ParseAny(dateStr)
	if dateErr != nil {
		return dateErr
	}
	if validation.IsValidateTitle(title) {
		e.Title = title
		e.StartAt = date
		e.Priority = Priority(priority)
	} else {
		return errors.New("неверный заголовок")
	}
	return nil
}

func (e *Event) AddReminder(message string, at time.Time) {
	e.Reminder, _ = reminder.NewReminder(message, at)
}

func (e *Event) RemoveReminder() {
	e.Reminder = nil
}

func NewEvent(title, dateStr, priority string) (*Event, error) {
	isValid := validation.IsValidateTitle(title)
	if isValid {
		dateParser, err := dateparse.ParseAny(dateStr)
		if err != nil {
			return nil, errors.New("неверный формат даты")
		}
		return &Event{
			ID:       getNextID(),
			Title:    title,
			StartAt:  dateParser,
			Priority: Priority(priority),
			Reminder: nil,
		}, nil
	}
	return &Event{}, errors.New("неправильный формат имени")
}
