package calendar

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/Kaktysfo/app/events"
	"github.com/Kaktysfo/app/storage"
)

var (
	EventError      = errors.New("cобытие с таким именем уже существует")
	addEventError   = errors.New("ошибка добавления события")
	deleteFindError = errors.New("событие с таким именем не найдено")
	showError       = errors.New("список событий пуст")
	saveError       = errors.New("ошибка маршлинга")
	deserError      = errors.New("ошибка десериализации")
	cSaveError      = errors.New("ошибка сохранения события")
	storageNilError = errors.New("хранилище пустое")
	EventRemindNil  = errors.New("у события с таким ID нет активного напоминания")
)

type Calendar struct {
	calendarEvents map[string]*events.Event
	storage        storage.Store
	Notification   chan string
}

func NewCalendar(s storage.Store) *Calendar {
	return &Calendar{
		calendarEvents: make(map[string]*events.Event),
		storage:        s,
		Notification:   make(chan string),
	}
}

func (c *Calendar) AddEvent(name string, date string, priority events.Priority) (*events.Event, error) {
	for _, event := range c.calendarEvents {
		if event.Title == name {
			return nil, EventError
		}
	}
	event, err := events.NewEvent(name, date, priority)
	if err != nil {
		return nil, addEventError
	}
	c.calendarEvents[event.ID] = event
	fmt.Println("Событие добавлено:", event.Title)
	if err := c.Save(); err != nil {
		return nil, cSaveError
	}
	return event, nil
}

func (c *Calendar) SetEventReminder(ID string, reminderMessage string, reminderTimer time.Time) error {
	err := c.isEventExistByID(ID)
	if err != nil {
		return err
	}
	event := c.calendarEvents[ID]
	event.AddReminder(reminderMessage, reminderTimer, c.Notify)
	return nil
}

func (c *Calendar) CancelEventReminder(ID string) error {
	err := c.isEventExistByID(ID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	event := c.calendarEvents[ID]
	if event.Reminder == nil {
		return EventRemindNil
	}
	event.RemoveReminder()
	err2 := c.Save()
	if err2 != nil {
		return err2
	}
	return nil
}

func (c *Calendar) Notify(msg string) {
	c.Notification <- msg
}

func (c *Calendar) ShowEvents() error {
	if len(c.calendarEvents) == 0 {
		return showError
	}
	fmt.Println("▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼")
	for _, event := range c.calendarEvents {
		fmt.Printf(
			"\nID события: %s || Название события: %s || Дата и время события: %s || Приоритет: %s",
			event.ID,
			event.Title,
			event.StartAt,
			event.Priority,
		)
		fmt.Println("")
	}
	fmt.Println("▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲")
	return nil
}

func (c *Calendar) isEventExistByID(id string) error {
	if _, ok := c.calendarEvents[id]; !ok {
		return EventError
	}
	return nil
}

func (c *Calendar) DeleteEvent(ID string) error {
	err := c.isEventExistByID(ID)
	if err != nil {
		return deleteFindError
	}
	delete(c.calendarEvents, ID)
	return nil
}

func (c *Calendar) EditEvent(id, newTitle, dateStr string, priority events.Priority) error {
	e, exists := c.calendarEvents[id]
	if !exists {
		return errors.New("не удалось найти событие")
	}
	err := e.Update(newTitle, dateStr, priority)
	if err != nil {
		return err
	}
	return c.Save()
}

func (c *Calendar) Save() error {
	if c.storage == nil {
		return storageNilError
	}
	data, err := json.Marshal(c.calendarEvents)
	if err != nil {
		return saveError
	}
	err2 := c.storage.Save(data)
	if err2 != nil {
		return err2
	}
	return nil
}

func (c *Calendar) Load() error {
	data, err := c.storage.Load()
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &c.calendarEvents)
	if err != nil {
		return deserError
	}
	return nil
}
