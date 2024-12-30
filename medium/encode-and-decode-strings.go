package medium

import (
	"fmt"
	"strings"
	"unicode"
)

func Encode(strs []string) string {
	var sb strings.Builder
	var fs string
	for _, s := range strs {
		fs = fmt.Sprintf("%d#%s", len(s), s)
		sb.WriteString(fs)
	}

	return sb.String()
}

func Decode(str string) []string {
	var fl []string = make([]string, 0)
	var i int = 0
	var digit int = 0
	for i < len(str) {
		digit = 0
		for unicode.IsDigit(rune(str[i])) {
			digit *= 10
			digit += int(str[i] - '0')
			i++
		}
		i++

		fl = append(fl, str[i:i+digit])

		i += digit
	}
	return fl
}
