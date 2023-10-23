package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

//	func main() {
//		s, err := Unpack("aasdф2ы\\3yй5s2qq0w")
//		//s, err := Unpack("a4bc2d5ef")
//		fmt.Println(s, err)
//	}
func Unpack(input string) (string, error) {
	// Place your code here.
	result := ""
	runes := []rune(input)
	for i := 0; i < len(runes); i++ {
		switch {
		// first+ rune expected not to have +1 and +2 positions as numbers
		case i+2 < len(runes) && unicode.IsDigit(runes[i+1]) && unicode.IsDigit(runes[i+2]):
			message := "Digits are not allowed, only numbers [0-9]:" + string(runes[i+1]) + string(runes[i+2])
			return message, ErrInvalidString

		// first+ rune is not a number, keep it as is
		case i+1 < len(runes) && !unicode.IsDigit(runes[i]) && !unicode.IsDigit(runes[i+1]):
			result += string(runes[i])

		// +0 and +1 runes are notDigit and isDigit, so repeat +0 rune +1 times
		case i+1 < len(runes) && !unicode.IsDigit(runes[i]) && unicode.IsDigit(runes[i+1]):
			times, _ := strconv.Atoi(string(runes[i+1]))
			result += strings.Repeat(string(runes[i]), times)
			i++ // skip rune with times itself

		// first rune not a number
		case i == 0:
			check, err := strconv.Atoi(string(runes[i]))
			if err == nil {
				message := "First rune cannot be a number:" + strconv.Itoa(check)
				return message, ErrInvalidString
			}

		// if last rune is not a number - keep it as is
		case i+1 == len(runes) && !unicode.IsDigit(runes[i]):
			result += string(runes[i])
		}
	}
	return result, nil
}
