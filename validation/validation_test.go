package validation

import (
	"testing"
)

func TestValidTitle(t *testing.T) {
	title := "привет"
	ok := IsValidateTitle(title)
	if !ok {
		t.Errorf("Invalid title validate %v", ok)
	}

	title = "hello67"
	ok = IsValidateTitle(title)
	if !ok {
		t.Errorf("Invalid title_number validate %v", ok)
	}
}
