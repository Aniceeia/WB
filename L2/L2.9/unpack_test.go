package unpack

import (
	"testing"
)

// тесты базовой функциональности
func TestUnpack_EmptyString(t *testing.T) {
	got, err := Unpack("")
	if err != nil {
		t.Errorf("unpack() error = %v, want nil", err)
	}
	if got != "" {
		t.Errorf("unpack() = %v, want empty string", got)
	}
}

func TestUnpack_NoNumbers(t *testing.T) {
	input := "abcd"
	want := "abcd"

	got, err := Unpack(input)
	if err != nil {
		t.Errorf("unpack() error = %v, want nil", err)
	}
	if got != want {
		t.Errorf("unpack() = %v, want %v", got, want)
	}
}

// тесты распаковки
func TestUnpack_SimpleUnpacking(t *testing.T) {
	input := "a4bc2d5e"
	want := "aaaabccddddde"

	got, err := Unpack(input)
	if err != nil {
		t.Errorf("unpack() error = %v, want nil", err)
	}
	if got != want {
		t.Errorf("unpack() = %v, want %v", got, want)
	}
}

func TestUnpack_MultiDigitNumber(t *testing.T) {
	input := "a10b"
	want := "aaaaaaaaaab"

	got, err := Unpack(input)
	if err != nil {
		t.Errorf("unpack() error = %v, want nil", err)
	}
	if got != want {
		t.Errorf("unpack() = %v, want %v", got, want)
	}
}

func TestUnpack_SpecialCharacters(t *testing.T) {
	input := "!2?3"
	want := "!!???"

	got, err := Unpack(input)
	if err != nil {
		t.Errorf("unpack() error = %v, want nil", err)
	}
	if got != want {
		t.Errorf("unpack() = %v, want %v", got, want)
	}
}

func TestUnpack_Spaces(t *testing.T) {
	input := "a b"
	want := "a b"

	got, err := Unpack(input)
	if err != nil {
		t.Errorf("unpack() error = %v, want nil", err)
	}
	if got != want {
		t.Errorf("unpack() = %v, want %v", got, want)
	}
}

func TestUnpack_SpacesWithNumbers(t *testing.T) {
	input := "a 3b"
	want := "a   b"

	got, err := Unpack(input)
	if err != nil {
		t.Errorf("unpack() error = %v, want nil", err)
	}
	if got != want {
		t.Errorf("unpack() = %v, want %v", got, want)
	}
}

// тесты обработки цифр после escape-символов
func TestUnpack_EscapedNumberFollowedByMultiplier(t *testing.T) {
	input := "\\45"
	want := "44444"

	got, err := Unpack(input)
	if err != nil {
		t.Errorf("unpack() error = %v, want nil", err)
	}
	if got != want {
		t.Errorf("unpack() = %v, want %v", got, want)
	}
}

func TestUnpack_LetterEscapedThenMultiplier(t *testing.T) {
	input := "a\\43"
	want := "a444"

	got, err := Unpack(input)
	if err != nil {
		t.Errorf("unpack() error = %v, want nil", err)
	}
	if got != want {
		t.Errorf("unpack() = %v, want %v", got, want)
	}
}

// тесты edge cases для isEndOfString
func TestUnpack_LastCharacterIsEscape(t *testing.T) {
	input := "abc\\"

	_, err := Unpack(input)
	if err == nil {
		t.Errorf("unpack() should return error for input %v", input)
	}
}

func TestUnpack_SingleEscapeCharacter(t *testing.T) {
	input := "\\"

	_, err := Unpack(input)
	if err == nil {
		t.Errorf("unpack() should return error for input %v", input)
	}
}

func TestUnpack_StringEndsWithLetter(t *testing.T) {
	input := "abc"
	want := "abc"

	got, err := Unpack(input)
	if err != nil {
		t.Errorf("unpack() error = %v, want nil", err)
	}
	if got != want {
		t.Errorf("unpack() = %v, want %v", got, want)
	}
}

// тесты двойного escape-символа
func TestUnpack_DoubleEscapeSymbol(t *testing.T) {
	input := "\\\\"
	want := "\\"

	got, err := Unpack(input)
	if err != nil {
		t.Errorf("unpack() error = %v, want nil", err)
	}
	if got != want {
		t.Errorf("unpack() = %v, want %v", got, want)
	}
}

func TestUnpack_EscapeBeforeEscape(t *testing.T) {
	input := "n\\32"
	want := "n33"

	got, err := Unpack(input)
	if err != nil {
		t.Errorf("unpack() error = %v, want nil", err)
	}
	if got != want {
		t.Errorf("unpack() = %v, want %v", got, want)
	}
}

func TestUnpack_MultipleEscapeSequences(t *testing.T) {
	input := "\\1\\2\\3"
	want := "123"

	got, err := Unpack(input)
	if err != nil {
		t.Errorf("unpack() error = %v, want nil", err)
	}
	if got != want {
		t.Errorf("unpack() = %v, want %v", got, want)
	}
}

func TestUnpack_EscapeThenEscapeThenMultiplier(t *testing.T) {
	input := "\\\\\\3"
	want := "\\3"

	got, err := Unpack(input)
	if err != nil {
		t.Errorf("unpack() error = %v, want nil", err)
	}
	if got != want {
		t.Errorf("unpack() = %v, want %v", got, want)
	}
}

// тесты out-of-bounds сценариев
func TestUnpack_SingleDigitAtEnd(t *testing.T) {
	input := "a"
	want := "a"

	got, err := Unpack(input)
	if err != nil {
		t.Errorf("unpack() error = %v, want nil", err)
	}
	if got != want {
		t.Errorf("unpack() = %v, want %v", got, want)
	}
}

func TestUnpack_DigitAtEndWithoutCharacter(t *testing.T) {
	// этот тест должен вернуть ошибку, так как начинается с цифры
	input := "4"

	_, err := Unpack(input)
	if err == nil {
		t.Errorf("unpack() should return error for input %v", input)
	}
}

func TestUnpack_CharacterWithMultiplierAtEnd(t *testing.T) {
	input := "a5"
	want := "aaaaa"

	got, err := Unpack(input)
	if err != nil {
		t.Errorf("unpack() error = %v, want nil", err)
	}
	if got != want {
		t.Errorf("unpack() = %v, want %v", got, want)
	}
}

// тесты ошибок
func TestUnpack_OnlyNumbers(t *testing.T) {
	input := "45"

	_, err := Unpack(input)
	if err == nil {
		t.Errorf("unpack() should return error for input %v", input)
	}
}

func TestUnpack_InvalidNumberPosition(t *testing.T) {
	input := "4a"

	_, err := Unpack(input)
	if err == nil {
		t.Errorf("unpack() should return error for input %v", input)
	}
}

func TestUnpack_NumberZero(t *testing.T) {
	input := "a0"

	_, err := Unpack(input)
	if err == nil {
		t.Errorf("unpack() should return error for input %v", input)
	}
}

// тесты escape-последовательностей
func TestUnpack_EscapedNumbers(t *testing.T) {
	input := "qwe\\4\\5"
	want := "qwe45"

	got, err := Unpack(input)
	if err != nil {
		t.Errorf("unpack() error = %v, want nil", err)
	}
	if got != want {
		t.Errorf("unpack() = %v, want %v", got, want)
	}
}

func TestUnpack_MixedEscapedAndUnpacked(t *testing.T) {
	input := "qwe\\45"
	want := "qwe44444"

	got, err := Unpack(input)
	if err != nil {
		t.Errorf("unpack() error = %v, want nil", err)
	}
	if got != want {
		t.Errorf("unpack() = %v, want %v", got, want)
	}
}

func TestUnpack_EscapedBackslash(t *testing.T) {
	input := "qwe\\\\5"
	want := "qwe\\\\\\\\\\"

	got, err := Unpack(input)
	if err != nil {
		t.Errorf("unpack() error = %v, want nil", err)
	}
	if got != want {
		t.Errorf("unpack() = %v, want %v", got, want)
	}
}

func TestUnpack_SingleEscapeAtEnd(t *testing.T) {
	input := "qwe\\"

	_, err := Unpack(input)
	if err == nil {
		t.Errorf("unpack() should return error for input %v", input)
	}
}

// дополнительные граничные тесты
func TestUnpack_SingleCharacter(t *testing.T) {
	input := "a"
	want := "a"

	got, err := Unpack(input)
	if err != nil {
		t.Errorf("unpack() error = %v, want nil", err)
	}
	if got != want {
		t.Errorf("unpack() = %v, want %v", got, want)
	}
}

func TestUnpack_RepeatedSingleCharacter(t *testing.T) {
	input := "a5"
	want := "aaaaa"

	got, err := Unpack(input)
	if err != nil {
		t.Errorf("unpack() error = %v, want nil", err)
	}
	if got != want {
		t.Errorf("unpack() = %v, want %v", got, want)
	}
}

func TestUnpack_MultipleDigitsAfterEscape(t *testing.T) {
	input := "\\410"
	want := "4444444444"

	got, err := Unpack(input)
	if err != nil {
		t.Errorf("unpack() error = %v, want nil", err)
	}
	if got != want {
		t.Errorf("unpack() = %v, want %v", got, want)
	}
}

func TestUnpack_EscapedCharacterRepeated(t *testing.T) {
	input := "\\43"
	want := "444"

	got, err := Unpack(input)
	if err != nil {
		t.Errorf("unpack() error = %v, want nil", err)
	}
	if got != want {
		t.Errorf("unpack() = %v, want %v", got, want)
	}
}

// тест для комплексного сценария
func TestUnpack_ComplexScenario(t *testing.T) {
	input := "a2\\3bcd\\4"
	want := "aa3bcd4"

	got, err := Unpack(input)
	if err != nil {
		t.Errorf("unpack() error = %v, want nil", err)
	}
	if got != want {
		t.Errorf("unpack() = %v, want %v", got, want)
	}
}
