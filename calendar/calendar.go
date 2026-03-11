package calendar

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Kaktysfo/app/events"
	"github.com/Kaktysfo/app/storage"
)

//var calendarEvents = make(map[string]*events.Event)

var (
	EventError      = errors.New("cобытие с таким именем уже существует")
	addEventError   = errors.New("ошибка добавления события")
	deleteError     = errors.New("ошибка удаления события")
	showError       = errors.New("список событий пуст")
	saveError       = errors.New("ошибка маршлинга")
	deserError      = errors.New("ошибка десериализации")
	cSaveError      = errors.New("ошибка сохранения события")
	storageNilError = errors.New("хранилище пустое")
)

type Calendar struct {
	calendarEvents map[string]*events.Event
	storage        *storage.Storage
}

func NewCalendar(s *storage.Storage) *Calendar {
	return &Calendar{
		calendarEvents: make(map[string]*events.Event),
		storage:        s,
	}
}

func (c *Calendar) AddEvent(name, date string) (*events.Event, error) {
	for _, event := range c.calendarEvents {
		if event.Title == name {
			return nil, EventError
		}
	}
	event, err := events.NewEvent(name, date)
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

func (c *Calendar) ShowEvents() error {
	if len(c.calendarEvents) == 0 {
		return showError
	}
	fmt.Println("▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼")
	for _, event := range c.calendarEvents {
		fmt.Printf(
			"\nНазвание события: %s || Дата и время события: %s ||",
			event.Title,
			event.StartAt,
		)
	}

	fmt.Println("\n▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲")

	return nil
}

func (c *Calendar) isEventExist(id string) error {
	if _, ok := c.calendarEvents[id]; !ok {
		return EventError
	}
	return nil
}

func (c *Calendar) DeleteEvent(name string) error {
	err := c.isEventExist(name)
	if err != nil {
		return deleteError
	}
	delete(c.calendarEvents, name)
	fmt.Println("Успешно удалено!", name)
	return nil
}

func (c *Calendar) EditEvent(id, newTitle, dateStr string) error {
	e, exists := c.calendarEvents[id]
	if !exists {
		return errors.New("не удалось найти событие")
	}
	err := e.Update(newTitle, dateStr)
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

//func fullValidation(name, title string) error {
//	if _, ok := calendarEvents[name]; !ok {
//		return TitleError
//	}
//	if ok := validation.IsValidateTitle(title); !ok {
//		return errors.New("введен некорректно заголовок")
//	}
//	if calendarEvents[name].Title == title {
//		return errors.New("такой заголовок уже существует")
//	}
//	return nil
//}
