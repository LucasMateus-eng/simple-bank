package validations

import "strings"

// StringIsEmptyOrWhiteSpace função responsável por verificar string completamente vazia
func StringIsEmptyOrWhiteSpace(str string) bool {
	return len([]rune(strings.TrimSpace(str))) == 0
}
