package validation

import (
	"errors"
	"fmt"
	"regexp"
)

var (
	errTimerStop = errors.New("таймер пуст или не запущен")
	titleErr     = errors.New("неверный формат заголовка")
)

func IsValidateTitle(title string) bool {
	pattern := "^[a-zA-Zа-яА-Я0-9 ,/.]{3,250}$"
	matched, err := regexp.MatchString(pattern, title)
	if err != nil {
		return false
	}
	return matched
}

func OpenError(err error) {
	switch {
	case errors.Is(err, errTimerStop):
		fmt.Println("Таймер не запущен, запустите его с помощью команды: remind")
	case errors.Is(err, titleErr):
		fmt.Println("Неверный формат сообщения")
	default:
		fmt.Println("Упс! Что-то пошло не так.")
		fmt.Println("")
	}

}
