package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var res strings.Builder
	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			return "", ErrInvalidString
		}

		if i+1 < len(s) && unicode.IsDigit(rune(s[i+1])) {
			repCnt, _ := strconv.Atoi(string(s[i+1]))
			res.WriteString(strings.Repeat(string(s[i]), repCnt))
			i++
		} else {
			res.WriteString(string(s[i]))
		}
	}

	return res.String(), nil
}
