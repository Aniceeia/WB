package unpack

import (
	"errors"
	"strings"
	"unicode"
)

// ErrInvalidString возвращается при некорректной входной строке
var ErrInvalidString = errors.New("invalid string")

// Unpack выполняет распаковку строки
func Unpack(input string) (string, error) {
	if isEmpty(input) {
		return "", nil
	}
	result := &strings.Builder{}
	runes := []rune(input)
	position := 0

	for position < len(runes) {
		char := runes[position]

		if isEscapeSymbol(char) {
			err := handleEscapedSymbol(result, runes, &position)
			if err != nil {
				return "", err
			}
			continue
		}

		if unicode.IsDigit(char) && position == 0 {
			return "", ErrInvalidString
		}

		err := processSymbol(result, runes, &position)
		if err != nil {
			return "", err
		}
	}

	return result.String(), nil
}

// проверяет, является ли строка пустой
func isEmpty(input string) bool {
	return input == ""
}

// проверяет, является ли символ экранирующим
func isEscapeSymbol(char rune) bool {
	return char == '\\'
}

// обрабатывает escape-последовательность
func handleEscapedSymbol(result *strings.Builder, runes []rune, position *int) error {
	if isEndOfString(runes, *position) {
		return ErrInvalidString
	}

	nextSymbol := runes[*position+1]

	if nextSymbol == '\\' || unicode.IsDigit(nextSymbol) {

		*position += 2
		return processEscapedSymbol(result, runes, position, nextSymbol)
	}

	return ErrInvalidString
}

// обрабатывает уже экранированный символ
func processEscapedSymbol(result *strings.Builder, runes []rune, position *int, nextSymbol rune) error {
	if *position < len(runes) && unicode.IsDigit(runes[*position]) {

		count, digitsCount := parseNumberFromRunes(runes[*position:])
		if count == 0 {
			return ErrInvalidString
		}

		repeatSymbol(result, nextSymbol, count)
		*position += digitsCount
	} else {
		result.WriteRune(nextSymbol)
	}

	return nil
}

// обрабатывает текущий символ
func processSymbol(result *strings.Builder, runes []rune, position *int) error {
	currentSymbol := runes[*position]

	if isEndOfString(runes, *position) || !unicode.IsDigit(runes[*position+1]) {
		result.WriteRune(currentSymbol)
		*position++
		return nil
	}

	count, digitsCount := parseNumberFromRunes(runes[*position+1:])
	if count == 0 {
		return ErrInvalidString
	}

	repeatSymbol(result, currentSymbol, count)
	*position += 1 + digitsCount

	return nil
}

// проверяет находится ли позиция в конце строки
func isEndOfString(runes []rune, position int) bool {
	return position+1 >= len(runes)
}

// извлекает число из начала среза рун
func parseNumberFromRunes(runes []rune) (int, int) {
	digitsCount := countDigits(runes)
	if digitsCount == 0 {
		return 0, 0
	}
	number := digitsToInteger(runes[:digitsCount])
	if number == 0 {
		return 0, digitsCount
	}
	return number, digitsCount
}

// подсчитывает количество цифр подряд
func countDigits(runes []rune) int {
	count := 0
	for count < len(runes) && unicode.IsDigit(runes[count]) {
		count++
	}
	return count
}

// преобразует срез цифр в число
func digitsToInteger(digits []rune) int {
	num := 0
	for _, digit := range digits {
		num = num*10 + int(digit-'0')
	}
	return num
}

// повторяет символ заданное количество раз
func repeatSymbol(result *strings.Builder, char rune, count int) {
	repeated := strings.Repeat(string(char), count)
	result.WriteString(repeated)
}
