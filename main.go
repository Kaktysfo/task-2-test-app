package main

import (
	"fmt"

	"github.com/Kaktysfo/app/calendar"
	"github.com/Kaktysfo/app/cmd"
	"github.com/Kaktysfo/app/logger"
	"github.com/Kaktysfo/app/storage"
)

func main() {
	fmt.Println("----Календарь----")
	logger.CreateLogger("app.log")
	defer logger.ExitLogger()
	logger.System("Запуск календаря")
	jsonStorage := storage.NewJsonStorage("calendar.json")
	logger.System("Создано хранилище")
	c := calendar.NewCalendar(jsonStorage)
	err := c.Load()
	if err != nil {
		fmt.Println("Ошибка загрузки:", err)
	}
	defer func() {
		errS := c.Save()
		if errS != nil {
			fmt.Println("Ошибка сохранения:", errS)
		}
	}()
	logger.System("Создание и загрузка календаря")
	cli := cmd.NewCmd(c)
	logger.System("Начало ввода")
	cli.Run()
}
