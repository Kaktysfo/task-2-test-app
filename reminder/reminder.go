package reminder

import (
	"errors"
	"fmt"
	"time"

	"github.com/Kaktysfo/app/validation"
)

type Reminder struct {
	Message  string
	RemindAt time.Time
	Sent     bool
}

func NewReminder(message string, at time.Time) (*Reminder, error) {
	check := validation.IsValidateTitle(message)
	if check == false {
		return nil, errors.New("неверный формат заголовка")
	}
	return &Reminder{
		Message:  message,
		RemindAt: at,
		Sent:     false,
	}, nil
}

func (r *Reminder) Send() {
	if r.Sent {
		return
	}
	fmt.Println("Напоминание!", r.Message)
	r.Sent = true
}

func (r *Reminder) Stop() {
}
