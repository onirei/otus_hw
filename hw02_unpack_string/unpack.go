package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var backslash bool
	var s strings.Builder
	var temp rune

	for _, char := range str {
		switch {
		case char == '\\' && !backslash:
			backslash = true
		case backslash && !unicode.IsDigit(char) && char != '\\':
			return "", ErrInvalidString
		case temp == 0 && (!unicode.IsDigit(char) || backslash):
			temp = char
			backslash = false
		case unicode.IsDigit(char) && !backslash:
			if temp == 0 {
				return "", ErrInvalidString
			}
			number, err := strconv.Atoi(string(char))
			if err == nil {
				s.WriteString(strings.Repeat(string(temp), number))
			}
			temp = 0
		default:
			if backslash && unicode.IsLetter(char) {
				return "", ErrInvalidString
			}
			s.WriteRune(temp)
			temp = char
			backslash = false
		}
	}
	if temp != 0 {
		s.WriteRune(temp)
	}
	fmt.Println(s.String())
	return s.String(), nil
}
