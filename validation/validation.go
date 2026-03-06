package validation

import (
	"regexp"
)

func IsValidateTitle(title string) bool {
	pattern := "^[a-zA-Zа-яА-Я0-9 ]{3,250}$"
	matched, err := regexp.MatchString(pattern, title)
	if err != nil {
		return false
	}
	return matched
}
