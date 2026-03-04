package calendar

import (
	"fmt"

	"github.com/Kaktysfo/app/events"
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
