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
	} else {
		calendar.AddEvent("event 1", event1)
	}
	event2, err2 := events.NewEvent("Уйти в запой", "03.09.2026")
	if err2 != nil {
		fmt.Println("Ошибка создания события:", err2)
	} else {
		calendar.AddEvent("event 2", event2)
	}
	event3, err3 := events.NewEvent("", "03.09.2026")
	if err3 != nil {
		fmt.Println("Ошибка создания события:", err3)
	} else {
		calendar.AddEvent("event 3", event3)
	}
	calendar.ShowEvents()
	_, errScan := fmt.Scanln()
	if errScan != nil {
		return
	}
}
