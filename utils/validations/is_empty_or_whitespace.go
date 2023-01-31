package validations

import "strings"

func StringIsEmptyOrWhiteSpace(str string) bool {
	return len([]rune(strings.TrimSpace(str))) == 0
}
