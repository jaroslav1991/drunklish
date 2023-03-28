package validator

import (
	"strings"
)

func ValidateDomain(email string) bool {
	validEmail := []string{"@yandex.ru", "@mail.ru", "@gmail.com", "@yahoo.com"}

	for _, mail := range validEmail {
		if strings.HasSuffix(email, mail) && len(email) > len(mail) {
			return true
		}
	}
	return false
}

func ValidateSymbol(email string) bool {
	var counter string

	for _, symbol := range email {
		if string(symbol) == string('@') {
			counter += string('@')
		}
	}

	if len(counter) == 1 {
		return true
	}

	return false
}

func LengthPassword(password string) bool {
	if len(password) > 5 {
		return true
	}
	return false
}
