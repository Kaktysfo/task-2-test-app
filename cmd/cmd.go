package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/Kaktysfo/app/calendar"
	"github.com/Kaktysfo/app/events"
	"github.com/c-bata/go-prompt"
	"github.com/google/shlex"
)

// var (
// 	parsInputErr = errors.New("ошибка парсинга ввода")
// )

type Cmd struct {
	calendar *calendar.Calendar
}

func NewCmd(c *calendar.Calendar) *Cmd {
	return &Cmd{
		calendar: c,
	}
}

func (c *Cmd) executor(input string) {
	err := c.calendar.Load()
	if err != nil {
		fmt.Println("Ошибка загрузки:", err)
	}
	parts, err := shlex.Split(input)
	if err != nil {
		fmt.Println("Ошибка парсинга ввода")
		return
	}
	cmd := strings.ToLower(parts[0])

	switch cmd {
	case "add":
		if len(parts) < 4 {
			fmt.Println("Формат: add \"название события\" \"дата и время\" \"приоритет\"")
			return
		}

		title := parts[1]
		date := parts[2]
		priority := events.Priority(parts[3])

		e, err := c.calendar.AddEvent(title, date, priority)
		if err != nil {
			fmt.Println("Ошибка добавления:", err)
		} else {
			fmt.Printf("Событие %s добавлено", e.Title)
		}
		err2 := c.calendar.Save()
		if err2 != nil {
			fmt.Println(err2)
			return
		}
	case "list":
		c.calendar.ShowEvents()
	case "help":
		fmt.Println("╔═══════════════════════════════════════╗")
		fmt.Println("║          ДОСТУПНЫЕ КОМАНДЫ            ║")
		fmt.Println("╠═══════════════════════════════════════╣")
		fmt.Println("║ list    → показать список задач       ║")
		fmt.Println("║ add     → добавить задачу             ║")
		fmt.Println("║ remove  → удалить задачу              ║")
		fmt.Println("║ search  → найти задачу                ║")
		fmt.Println("║ exit    → завершить программу         ║")
		fmt.Println("╚═══════════════════════════════════════╝")

	case "exit":
		err := c.calendar.Save()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Пока-пока")
		os.Exit(0)
	default:
		fmt.Println("Неизвестная команда:")
		fmt.Println("Введите 'help' для получения списка команд")
	}
}
func (c *Cmd) completer(d prompt.Document) []prompt.Suggest {
	suggestions := []prompt.Suggest{
		{Text: "add", Description: "Добавить событие"},
		{Text: "remove", Description: "Удалить событие"},
		{Text: "list", Description: "Показать все события"},
		{Text: "help", Description: "Показать справку"},
		{Text: "exit", Description: "Выйти из программы"},
	}
	return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
}

func (c *Cmd) Run() {
	p := prompt.New(
		c.executor,
		c.completer,
		prompt.OptionPrefix("> "),
	)
	p.Run()
}
