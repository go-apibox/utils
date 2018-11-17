package utils

import (
	"strings"
)

func CamelCase(str string) string {
	newstr := make([]rune, 0)
	upNextChar := true

	str = strings.ToLower(str)

	for _, chr := range str {
		switch {
		case upNextChar:
			upNextChar = false
			if 'a' <= chr && chr <= 'z' {
				chr -= ('a' - 'A')
			}
		case chr == '_':
			upNextChar = true
			continue
		}

		newstr = append(newstr, chr)
	}

	return string(newstr)
}

func SnakeCase(str string) string {
	newstr := make([]rune, 0)
	for idx, chr := range str {
		if isUpper := 'A' <= chr && chr <= 'Z'; isUpper {
			if idx > 0 {
				newstr = append(newstr, '_')
			}
			chr -= ('A' - 'a')
		}
		newstr = append(newstr, chr)
	}

	return string(newstr)
}
