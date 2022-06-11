package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var b strings.Builder
	digitFlag := false
	var prevElem rune
	counter := 1

	for indx, elem := range str {
		if unicode.IsDigit(elem) {
			if indx == 0 || digitFlag {
				return "", ErrInvalidString
			}

			counter, _ = strconv.Atoi(string(elem))
			if counter > 0 {
				b.WriteString(strings.Repeat(string(prevElem), counter-1))
			}
			digitFlag = true
		} else {
			if counter > 0 {
				b.WriteRune(prevElem)
			}
			prevElem = elem
			counter = 1
			digitFlag = false
		}
	}

	if counter > 0 {
		b.WriteRune(prevElem)
	}

	return b.String()[1:], nil
}
