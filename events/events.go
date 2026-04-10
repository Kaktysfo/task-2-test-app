package events

import (
	"errors"
	"fmt"
	"time"

	"github.com/Kaktysfo/app/logger"
	"github.com/Kaktysfo/app/reminder"
	"github.com/Kaktysfo/app/validation"
	"github.com/araddon/dateparse"
	"github.com/google/uuid"
)

const (
	sucDeleted = "Напоминание успешно удалено!"
	sendErr    = "Не удалось удалить напоминание: оно уже отправляется."
)

type Event struct {
	ID       string             `json:"ID"`
	Title    string             `json:"Title"`
	StartAt  time.Time          `json:"Time"`
	Priority Priority           `json:"Priority"`
	Reminder *reminder.Reminder `json:"Reminder"`
}

func getNextID() string {
	return uuid.New().String()
}

func PrepareData(title, dateStr string, priority Priority) (time.Time, error) {
	date, dateErr := dateparse.ParseAny(dateStr)
	if dateErr != nil {
		return time.Time{}, dateErr
	}
	if !validation.IsValidateTitle(title) {
		return time.Time{}, errors.New("неверный заголовок")
	}
	return date, nil
}

func (e *Event) Update(title, dateStr string, priority Priority) error {
	parsedDate, err := PrepareData(title, dateStr, priority)
	if err != nil {
		return err
	}
	e.Title = title
	e.StartAt = parsedDate
	e.Priority = priority
	return nil
}

func (e *Event) AddReminder(message string, timerAT time.Time, notify func(msg string)) error {
	r, err := reminder.NewReminder(message, timerAT, notify)
	if err != nil {
		return err
	}
	if e.Reminder != nil {
		e.Reminder.Stop()
	}
	e.Reminder = r
	fmt.Printf("Новое напоминание установлено: %s на %v\n", message, timerAT)
	logger.Info("Установлено новое напоминание!")
	e.Reminder.Start()
	return nil
}

func (e *Event) RemoveReminder() {
	if e.Reminder == nil {
		fmt.Println("Напоминание не установлено.")
	}
	stopped, err := e.Reminder.Stop()
	if err != nil {
		validation.OpenError(err)
		return
	}
	if stopped {
		e.Reminder = nil
		fmt.Println(sucDeleted)
		logger.Info(sucDeleted)
	} else {
		fmt.Println(sendErr)
		logger.Info(sendErr)
	}
}

func NewEvent(title, dateStr string, p Priority) (*Event, error) {
	parsedDate, err := PrepareData(title, dateStr, p)
	if err != nil {
		return nil, err
	}
	return &Event{
		ID:       getNextID(),
		Title:    title,
		StartAt:  parsedDate,
		Priority: p,
		Reminder: nil,
	}, nil
}
