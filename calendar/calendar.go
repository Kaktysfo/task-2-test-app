package calendar

import (
	"errors"
	"fmt"

	"github.com/Kaktysfo/app/events"
)

var calendarEvents = make(map[string]*events.Event)
var TitleError = errors.New("cобытие с таким именем уже существует")
var addEventError = errors.New("ошибка добавления события")
var deleteError = errors.New("ошибка удаления события")
var showError = errors.New("список событий пуст")

type Calendar struct {
	calendarEvents map[string]*events.Event
}

func NewCalendar() *Calendar {
	return &Calendar{calendarEvents: calendarEvents}
}

func AddEvent(name, date string) (*events.Event, error) {
	event, err := events.NewEvent(name, date)
	if err != nil {
		return &events.Event{}, addEventError
	}
	calendarEvents[event.ID] = event
	fmt.Println("Событие добавлено: ", event.Title)
	return event, nil
}

func ShowEvents() error {
	if len(calendarEvents) == 0 {
		return showError
	}
	fmt.Println("\nВсе события в календаре: ")
	fmt.Println("▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼")
	for _, event := range calendarEvents {
		fmt.Printf("\nНазвание события:  %s || Дата и время события:  %s ||", event.Title, event.StartAt)
	}
	fmt.Println("\n \n▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲")
	return nil
}

func isTitleExist(title string) error {
	if _, ok := calendarEvents[title]; !ok {
		return TitleError
	}
	return nil
}

func DeleteEvent(name string) error {
	err := isTitleExist(name)
	if err != nil {
		return deleteError
	}
	delete(calendarEvents, name)
	fmt.Println("Успешно удалено!", name)
	return nil
}

func EditEvent(id, newTitle, dateStr string) error {
	e, exists := calendarEvents[id]
	if !exists {
		return errors.New("не удалось найти событие")
	}
	err := e.Update(newTitle, dateStr)
	return err
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
