package main

import (
	"fmt"

	"github.com/Kaktysfo/app/calendar"
	"github.com/Kaktysfo/app/cmd"
	"github.com/Kaktysfo/app/storage"
)

func main() {
	fmt.Println("----Календарь----")
	jsonStorage := storage.NewJsonStorage("calendar.json")
	c := calendar.NewCalendar(jsonStorage)

	// err := c.Load()
	// if err != nil {
	// 	fmt.Println("Ошибка загрузки:", err)
	// }
	cli := cmd.NewCmd(c)
	cli.Run()
	// event1, err1 := c.AddEvent("Завод", "2025/03/11", "Высокий")
	// if err1 != nil {
	// 	fmt.Println("Ошибка добавления события:", err1)
	// } else {
	// 	fmt.Println("Добавлено:", event1.Title)
	// }
	// event2, err2 := c.AddEvent("Работа", "2023/01/11", "Низкий")
	// if err2 != nil {
	// 	fmt.Println("Ошибка добавления события:", err2)
	// } else {
	// 	fmt.Println("Добавлено:", event2.Title)
	// }
	// event3, err3 := c.AddEvent("Кушац", "1999/12/12", "Средний")
	// if err3 != nil {
	// 	fmt.Println("Ошибка добавления события:", err3)
	// } else {
	// 	fmt.Println("Добавлено:", event3.Title)
	// }
	// err12 := c.EditEvent(event1.ID, "Завод UPD", "2000/02/02", "Низкий")
	// if err12 != nil {
	// 	fmt.Println("Ошибка редактирования:", err12)
	// }
	// err13 := c.DeleteEvent(event2.ID)
	// if err13 != nil {
	// 	fmt.Println("Ошибка удаления:", err13)
	// }
	c.ShowEvents()
	errS := c.Save()
	if errS != nil {
		fmt.Println("Ошибка сохранения:", errS)
	}
	fmt.Println(jsonStorage.GetFilename())
	// fmt.Println("----Новый календарь----")
	// zipStorage := storage.NewZipStorage("calendar.zip")
	// cz := calendar.NewCalendar(zipStorage)
	// err14 := cz.Load()
	// if err != nil {
	// 	fmt.Println("Ошибка загрузки:", err14)
	// }
	// zEvent1, err21 := cz.AddEvent("Тренировка", "2025/07/01 18:00", "Высокий")
	// if err21 != nil {
	// 	fmt.Println("Ошибка добавления события:", err21)
	// } else {
	// 	fmt.Println("Добавлено:", zEvent1.Title)
	// }
	// zipEvent2, err22 := cz.AddEvent("Шарага", "2025/07/02 09:30", "Средний")
	// if err22 != nil {
	// 	fmt.Println("Ошибка добавления события:", err22)
	// } else {
	// 	fmt.Println("Добавлено:", zipEvent2.Title)
	// }
	// err23 := cz.EditEvent(zEvent1.ID, "Созвон UPD", "2025/07/01 19:00", "Низкий")
	// if err23 != nil {
	// 	fmt.Println("Ошибка редактирования:", err23)
	// }
	// cz.ShowEvents()
	// err24 := cz.Save()
	// if err24 != nil {
	// 	fmt.Println("Ошибка сохранения:", err24)
	// }
	// fmt.Println(zipStorage.GetFilename())
}
