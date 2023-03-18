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
	sRunes := []rune(s)
	for i := 0; i < len(sRunes); i++ {
		if unicode.IsDigit(sRunes[i]) {
			return "", ErrInvalidString
		}

		if i+1 < len(sRunes) && unicode.IsDigit(sRunes[i+1]) {
			repCnt, err := strconv.Atoi(string(sRunes[i+1]))
			if err != nil {
				return "", err
			}
			res.WriteString(strings.Repeat(string(sRunes[i]), repCnt))
			i++
		} else {
			res.WriteString(string(sRunes[i]))
		}
	}

	return res.String(), nil
}
