package reminder

import (
	"errors"
	"fmt"
	"time"

	"github.com/Kaktysfo/app/logger"
	"github.com/Kaktysfo/app/validation"
)

const (
	remindWay  = "reminder/reminder.go: %w"
	remindSend = "Напоминание уже было отправлено."
)

var (
	errTimerStop = errors.New("таймер пуст или не запущен")
	titleErr     = errors.New("неверный формат заголовка")
)

type Reminder struct {
	Message  string
	RemindAt time.Time
	Sent     bool
	timer    *time.Timer
	notify   func(msg string)
}

func NewReminder(message string, at time.Time, notify func(msg string)) (*Reminder, error) {
	logger.System("Создание нового напоминания")
	check := validation.IsValidateTitle(message)
	if !check {
		return nil, fmt.Errorf(remindWay, titleErr)
	}
	return &Reminder{
		Message:  message,
		RemindAt: at,
		Sent:     false,
		timer:    nil,
		notify:   notify,
	}, nil
}

func (r *Reminder) Start() {
	logger.System("Запуск таймера.")
	now := time.Now()
	if r.timer != nil {
		r.timer.Stop()
		r.timer = nil
	}
	if r.RemindAt.Before(now) || r.RemindAt.Equal(now) {
		r.timer = time.AfterFunc(1*time.Millisecond, func() {
			r.Send()
		})
		return
	}
	duration := r.RemindAt.Sub(now)
	fmt.Println("Таймер установлен на", duration)
	logger.Info("Таймер установлен!")
	r.timer = time.AfterFunc(duration, func() {
		r.Send()
	})

}

func (r *Reminder) Send() {
	if r.Sent {
		fmt.Println(remindSend)
		logger.Info(remindSend)
		return
	}
	r.notify(r.Message)
	r.Sent = true
}

func (r *Reminder) Stop() (bool, error) {
	if r.timer == nil {
		return false, fmt.Errorf(remindWay, errTimerStop)
	}
	stopped := r.timer.Stop()
	r.timer = nil
	return stopped, nil
}
