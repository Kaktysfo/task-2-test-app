package main

import (
	"fmt"

	"github.com/Kaktysfo/app/calendar"
)

func main() {
	event1, err1 := calendar.AddEvent("День рождения", "2026/09/03 13:33")
	if err1 != nil {
		fmt.Println("Ошибка:", err1)
		return
	}
	event2, err2 := calendar.AddEvent("Уйти в запой", "03/09/2026")
	if err2 != nil {
		fmt.Println("Ошибка:", err2)
		return
	}
	errShow := calendar.ShowEvents()
	if errShow != nil {
		fmt.Println(errShow)
	}
	errDel := calendar.DeleteEvent(event1.ID)
	if errDel != nil {
		fmt.Println(errDel)
	}
	err := calendar.EditEvent(event2.ID, "Созвон", "2025/06/12 16:50")
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
	errShow2 := calendar.ShowEvents()
	if errShow2 != nil {
		fmt.Println(errShow2)
	}
}
