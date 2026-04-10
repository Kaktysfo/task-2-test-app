package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Kaktysfo/app/calendar"
	"github.com/Kaktysfo/app/events"
	"github.com/Kaktysfo/app/logger"
	"github.com/c-bata/go-prompt"
	"github.com/google/shlex"
)

const (
	eventUpd         = "Событие успешно изменено!"
	eventDeleted     = "Событие успешно удалено!"
	unknownCommand   = "Неизвестная команда: \n Введите 'help' для получения списка команд"
	removeFormat     = "Формат: remove \"ID события\""
	parsErr          = "Ошибка парсинга ввода"
	ErrOccured       = "Ошибка:"
	WrongTypeFormat  = "Неверный формат команды"
	EventAdded       = "Событие успешно добавлено!"
	AddErr           = "Ошибка добавления: "
	AddFormat        = "Формат команды: add \"название события\" \"дата и время\" \"приоритет\""
	ListFormat       = "Формат команды: \"list\""
	ExitFormat       = "Формат команды: \"exit\""
	UpdateFormat     = "Формат команды:\"ID события\" \"название события\" \"дата и время\" \"приоритет\""
	HelpFormat       = "Формат команды: help"
	RemindFormat     = "Формат команды: \"ID события\" \"сообщение напоминания\" \"время таймера\" "
	DelRemindFormat  = "Формат команды: \"ID события\" "
	LogFormat        = "Формат команды: log"
	timeTimerFormat  = "Форматы времени: 30s - 30 секунд | 10m - 10 минут | 2h- 2 часа | 1d - 1 день | 1h30m - 1 час 30 минут"
	parsingTypeError = "Ошибка парсинга ввода"
	nilType          = "Ввод не может быть пустым! Список команд: list"
)

type Cmd struct {
	calendar *calendar.Calendar
	logger   *Logs
}

func NewCmd(c *calendar.Calendar) *Cmd {
	return &Cmd{
		calendar: c,
		logger:   new(Logs),
	}
}

func (c *Cmd) executor(input string) {
	parts, err := shlex.Split(input)
	logger.Info(fmt.Sprintf("Пользовательский ввод: %v", input))
	if err != nil {
		fmt.Println(parsingTypeError)
		logger.Error(parsingTypeError)
		c.logger.Log(parsingTypeError, err)
		return
	}
	switch input {
	case "":
		fmt.Println(nilType)
		c.logger.Log(nilType, nil)
		logger.Error(nilType)
		return
	}
	cmd := strings.ToLower(parts[0])
	switch cmd {
	case "add":
		if len(parts) != 4 {
			fmt.Println(WrongTypeFormat)
			fmt.Println(AddFormat)
			c.logger.Log(WrongTypeFormat, nil)
			c.logger.Log(AddFormat, nil)
			logger.Info(WrongTypeFormat)
			logger.Info(AddFormat)
			return
		}
		title := parts[1]
		date := parts[2]
		priority := events.Priority(parts[3])
		err := events.Priority.Validate(priority)
		if err != nil {
			fmt.Println(ErrOccured)
			fmt.Println(err)
			c.logger.Log(ErrOccured, err)
			logger.Error(fmt.Sprintf(ErrOccured, err))
			return
		}
		e, err := c.calendar.AddEvent(title, date, priority)
		if err != nil {
			fmt.Println(AddErr, err)
			c.logger.Log(AddErr, err)
		} else {
			fmt.Println(EventAdded, e.Title)
			fmt.Println("")
			c.logger.Log(EventAdded, nil)
			logger.Info(EventAdded)
		}
	case "remove":
		if len(parts) != 2 {
			fmt.Println(WrongTypeFormat)
			fmt.Println(removeFormat)
			c.logger.Log(WrongTypeFormat, nil)
			c.logger.Log(removeFormat, nil)
			logger.Error(WrongTypeFormat)
			logger.Error(removeFormat)
			return
		}
		ID := parts[1]
		err := c.calendar.DeleteEvent(ID)
		if err != nil {
			fmt.Println(ErrOccured, err)
			c.logger.Log(ErrOccured, err)
			logger.Error(fmt.Sprintf(ErrOccured, err))
			return
		}
		fmt.Println(eventDeleted)
		fmt.Println("")
		c.logger.Log(eventDeleted, nil)
		logger.Info(eventDeleted)
	case "update":
		if len(parts) != 5 {
			fmt.Println(WrongTypeFormat)
			fmt.Println(UpdateFormat)
			c.logger.Log(WrongTypeFormat, nil)
			c.logger.Log(UpdateFormat, nil)
			logger.Info(WrongTypeFormat)
			logger.Info(UpdateFormat)
			return
		}
		ID := parts[1]
		title := parts[2]
		date := parts[3]
		priority := events.Priority(parts[4])
		errP := events.Priority.Validate(priority)
		if errP != nil {
			fmt.Println(ErrOccured)
			fmt.Println(errP)
			c.logger.Log(ErrOccured, errP)
			logger.Error(fmt.Sprintf(ErrOccured, errP))
			return
		}
		err := c.calendar.EditEvent(ID, title, date, priority)
		if err != nil {
			fmt.Println(ErrOccured, err)
			c.logger.Log(ErrOccured, err)
			logger.Error(fmt.Sprintf(ErrOccured, err))
			return
		}
		fmt.Println(eventUpd)
		fmt.Println("")
		c.logger.Log(eventUpd, nil)
		logger.Info(eventUpd)
	case "remind":
		if len(parts) != 4 {
			fmt.Println(WrongTypeFormat)
			fmt.Println(RemindFormat)
			fmt.Println(timeTimerFormat)
			c.logger.Log(WrongTypeFormat, nil)
			c.logger.Log(RemindFormat, nil)
			c.logger.Log(timeTimerFormat, nil)
			logger.Info(WrongTypeFormat)
			logger.Info(RemindFormat)
			logger.Info(timeTimerFormat)
			return
		}
		ID := parts[1]
		message := parts[2]
		timer := parts[3]
		duration, err2 := time.ParseDuration(timer)
		if err2 != nil {
			fmt.Println("Неверный формат длительности", err2)
			c.logger.Log(ErrOccured, err2)
			logger.Error(fmt.Sprintf(ErrOccured, err2))
			return
		}
		reminderTime := time.Now().Add(duration)
		err3 := c.calendar.SetEventReminder(ID, message, reminderTime)
		if err3 != nil {
			fmt.Println(ErrOccured, err3)
			c.logger.Log(ErrOccured, err3)
			logger.Error(fmt.Sprintf(ErrOccured, err3))
			return
		}
		fmt.Println("")
	case "delremind":
		if len(parts) != 2 {
			fmt.Println(WrongTypeFormat)
			fmt.Println(DelRemindFormat)
			c.logger.Log(WrongTypeFormat, nil)
			c.logger.Log(DelRemindFormat, nil)
			logger.Info(WrongTypeFormat)
			logger.Info(DelRemindFormat)
			return
		}
		ID := parts[1]
		err := c.calendar.CancelEventReminder(ID)
		if err != nil {
			fmt.Println(ErrOccured, err)
			c.logger.Log(ErrOccured, err)
			logger.Error(fmt.Sprintf(ErrOccured, err))
			return
		}
		fmt.Println("")
	case "log":
		if len(parts) != 1 {
			fmt.Println(WrongTypeFormat)
			fmt.Println(LogFormat)
			c.logger.Log(WrongTypeFormat, nil)
			c.logger.Log(LogFormat, nil)
			logger.Info(WrongTypeFormat)
			logger.Info(LogFormat)
			return
		}
		c.logger.ListLogger()
	case "list":
		if len(parts) != 1 {
			fmt.Println(WrongTypeFormat)
			fmt.Println(ListFormat)
			c.logger.Log(WrongTypeFormat, nil)
			c.logger.Log(ListFormat, nil)
			logger.Info(WrongTypeFormat)
			logger.Info(ListFormat)
			return
		}
		c.calendar.ShowEvents()
	case "help":
		if len(parts) != 1 {
			fmt.Println(WrongTypeFormat)
			fmt.Println(HelpFormat)
			c.logger.Log(WrongTypeFormat, nil)
			c.logger.Log(HelpFormat, nil)
			logger.Info(WrongTypeFormat)
			logger.Info(HelpFormat)
			return
		}
		fmt.Println("\n═══════════════════════════════════════════════════════════")
		fmt.Println("                      ДОСТУПНЫЕ КОМАНДЫ")
		fmt.Println("═══════════════════════════════════════════════════════════")
		fmt.Println(" list → показать все события                                                 ")
		fmt.Println("")
		fmt.Println(" add \"название\" \"дата время\" \"приоритет\" → добавить событие            ")
		fmt.Println("   Пример: add \"Встреча\" \"2024-12-25 15:30\" \"Высокий\"                     ")
		fmt.Println("   Приоритет: Низкий, Средний, Высокий                                              ")
		fmt.Println("                                                                             ")
		fmt.Println(" remove \"ID\"   → удалить событие                                           ")
		fmt.Println("   Пример: remove \"ID события\"                                         ")
		fmt.Println("                                                                             ")
		fmt.Println(" update \"ID\" \"название\" \"дата время\" \"приоритет\" → редактировать     ")
		fmt.Println("   Пример: update \"ID события\" \"Новая\" \"2024-12-26 10:00\" \"med\"  ")
		fmt.Println("                                                                             ")
		fmt.Println(" remind \"ID\" \"сообщение\" \"длительность\"   → создать напоминание        ")
		fmt.Println("   Пример: remind \"ID события\" \"Позвонить\" \"1h30m\"                    ")
		fmt.Println("   Форматы: 30s, 10m, 2h, 1d, 1h30m                                          ")
		fmt.Println("                                                                             ")
		fmt.Println(" delremind  \"ID\" → удаляет напоминание                                     ")
		fmt.Println("   Пример: delremind \"ID события\"   ")
		fmt.Println("")
		fmt.Println(" log    → показать логи  ")
		fmt.Println("")
		fmt.Println(" help   → показать доступные команд                                          ")
		fmt.Println("")
		fmt.Println(" exit   → завершить программу                                                ")
		fmt.Println("═══════════════════════════════════════════════════════════")
	case "exit":
		if len(parts) != 1 {
			fmt.Println(WrongTypeFormat)
			fmt.Println(ExitFormat)
			c.logger.Log(WrongTypeFormat, nil)
			c.logger.Log(ExitFormat, nil)
			logger.Info(WrongTypeFormat)
			logger.Info(ExitFormat)
			return
		}
		logger.Info("Запуск функции Save перед завершешнием программы")
		err := c.calendar.Save()
		if err != nil {
			fmt.Println(err)
			c.logger.Log(ErrOccured, err)
			logger.Error(fmt.Sprintf(ErrOccured, err))
			return
		}
		fmt.Println("Пока-пока")
		c.logger.Log("Пока-пока", nil)
		logger.System("Работа приложения завершена")
		close(c.calendar.Notification)
		os.Exit(0)
	default:
		fmt.Println(unknownCommand)
		c.logger.Log(unknownCommand, nil)
		logger.Info(unknownCommand)
	}
}
func (c *Cmd) completer(d prompt.Document) []prompt.Suggest {
	suggestions := []prompt.Suggest{
		{Text: "add", Description: "Добавить событие"},
		{Text: "remove", Description: "Удалить событие"},
		{Text: "update", Description: "Отредактировать событие"},
		{Text: "remind", Description: "Создать напоминание к задаче"},
		{Text: "delremind", Description: "Удалить напоминание к задаче"},
		{Text: "list", Description: "Показать все события"},
		{Text: "log", Description: "Показать все логи"},
		{Text: "help", Description: "Показать справку"},
		{Text: "exit", Description: "Выйти из программы"},
	}
	return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
}

func (c *Cmd) Run() {
	go func() {
		for msg := range c.calendar.Notification {
			fmt.Println(msg)
		}
	}()
	p := prompt.New(
		c.executor,
		c.completer,
		prompt.OptionPrefix("> "),
	)
	p.Run()
}
