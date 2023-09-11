package validation

import (
	"regexp"
	"unicode"
)

func ValidateEmail(email string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	return re.MatchString(email)
}

func VerifyPassword(input string) (sevenOrMore, number, upper, special bool) {
	letters := 0
	for _, input := range input {
		switch {
		case unicode.IsNumber(input):
			number = true
		case unicode.IsUpper(input):
			upper = true
			letters++
		case unicode.IsPunct(input) || unicode.IsSymbol(input):
			special = true
		case unicode.IsLetter(input) || input == ' ':
			letters++
		default:
			return false, false, false, false
		}
	}
	sevenOrMore = letters >= 7
	return
}
