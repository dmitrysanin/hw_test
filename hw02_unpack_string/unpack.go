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
	var digitFlag bool = false
	var prev_elem rune
	var counter int = 1

	for indx, elem := range str {

		if unicode.IsDigit(elem) {
			if indx == 0 || digitFlag {
				return "", ErrInvalidString
			}

			counter, _ = strconv.Atoi(string(elem))
			if counter > 0 {
				b.WriteString(strings.Repeat(string(prev_elem), counter-1))
			}
			digitFlag = true
		} else {
			if counter > 0 {
				b.WriteRune(prev_elem)
			}
			prev_elem = elem
			counter = 1
			digitFlag = false
		}
	}

	if counter > 0 {
		b.WriteRune(prev_elem)
	}

	return b.String()[1:], nil
}
