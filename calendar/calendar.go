package calendar

import (
	"fmt"

	"github.com/Kaktysfo/app/events"
	"github.com/araddon/dateparse"
)

var eventsMap = make(map[string]events.Event)

func AddEvent(key string, e events.Event) {
	eventsMap[key] = e
	fmt.Println("Событие добавлено: ", e.Title)
}

func ShowEvents() {
	fmt.Println("\nВсе события в календаре: ")
	fmt.Println("▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼")
	for _, events := range eventsMap {
		fmt.Printf("\nНазвание события:  %s || Дата и время события:  %s ||", events.Title, events.StartAt)
	}
	fmt.Println("\n \n▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲")
}

func DeleteEvent(key string) {
	delete(eventsMap, key)
	fmt.Println("Успешно удалено!", key)
}

func EditEvent(key string, newTitle string, date string) {
	for _, events := range eventsMap {
		if key == events.Title {
			events.Title = newTitle
			events.StartAt, err = dateparse.ParseAny(date)
			if err {

			}
		}
	}
}
