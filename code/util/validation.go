package util

import (
	"monitor/model"
	"strings"
	"unicode"
)

func ValidateUser(s model.RegisterUserRequest) bool {
	if !ValidatePassword(*s.PassWord) || !ValidateUserName(*s.UserName) {
		return false
	}
	return true
}

func ValidateUserName(u string) bool {
	if len(u) < 5 {
		return false
	}
	if !unicode.IsLetter(rune(u[0])) {
		return false
	}
	for _, v := range u {
		if !unicode.IsLetter(v) && !unicode.IsDigit(v) {
			return false
		}
	}
	return true
}

func ValidatePassword(u string) bool {
	has_digit := false
	has_lower := false
	has_upper := false
	has_charachter := false
	if len(u) < 10 {
		return false
	}
	for _, v := range u {
		if unicode.IsDigit(v) {
			has_digit = true
		} else if unicode.IsLower(v) {
			has_lower = true
		} else if unicode.IsUpper(v) {
			has_upper = true
		} else {
			has_charachter = true
		}
	}
	if !has_digit || !has_charachter || !has_lower || !has_upper {
		return false
	}
	return true
}

func ValidateMethod(u string) bool {
	m := strings.ToUpper(u)
	if m == "OPTIONS" || m == "HEAD" || m == "GET" {
		return true
	}
	return false
}
