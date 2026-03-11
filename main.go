package main

import (
	"fmt"

	"github.com/Kaktysfo/app/calendar"
)

func main() {
	c := calendar.NewCalendar()

	event1, err1 := calendar.AddEvent("Meeting", "2025/06/12")
	if err1 != nil {
		fmt.Println("Error:", err1)
	} else {
		fmt.Println(event1.Title, "added")
	}
	event2, err2 := calendar.AddEvent("One more meeting", "2025/06/12")
	if err2 != nil {
		fmt.Println("Error:", err2)
	} else {
		fmt.Println(event2.Title, "added")
	}
	err := calendar.EditEvent(event2.ID, "Call", "2025/06/12 16:50")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Event updated")
	}
	errShow2 := calendar.ShowEvents()
	if errShow2 != nil {
		fmt.Println(errShow2)
	}
	fmt.Scanln()
}
