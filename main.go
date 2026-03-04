package main

import (
	"fmt"

	"github.com/Kaktysfo/app/calendar"
	"github.com/Kaktysfo/app/events"
)

func main() {
	event1, err := events.NewEvent("День рождения", "2026/09/03 13:33")
	if err != nil {
		fmt.Println("Ошибка создания события:", err)
	}
	event2, err := events.NewEvent("Уйти в запой", "03.09.2026")
	calendar.AddEvent("event 1", event1)
	calendar.AddEvent("event 2", event2)
	calendar.ShowEvents()
	fmt.Scanln()
}
