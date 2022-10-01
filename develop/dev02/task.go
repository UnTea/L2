package main

import (
	"errors"
	"fmt"
	"os"
)

/*
	Задача на распаковку

	Создать Go-функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны, например:
	* "a4bc2d5e" => "aaaabccddddde"
	* "abcd" => "abcd"
	* "45" => "" (некорректная строка)
	* "" => ""

	Дополнительно

	Реализовать поддержку escape-последовательностей, например:
	* qwe\4\5 => qwe45 (*)
	* qwe\45 => qwe44444 (*)
	* qwe\\5 => qwe\\\\\ (*)

	В случае если была передана некорректная строка, функция должна возвращать ошибку.
	Написать unit-тесты.
*/

var errCorrectness = errors.New("the string must not start with a digit and contain two digits in a row")

// IsNumber is a function that returns true if character value is between '1' - '9'
func IsNumber(character rune) bool {
	if character >= '1' && character <= '9' {
		return true
	}

	return false
}

// IsSlash is a function that returns true if character is a slash
func IsSlash(character rune) bool {
	return character == '\\'
}

// IsCorrect is a function that checking for string correctness
func IsCorrect(r []rune) bool {
	if IsNumber(r[0]) {
		return false
	}

	var (
		twoDigit bool
		slash    bool
	)

	for i := range r {
		if IsSlash(r[i]) {
			slash = true

			continue
		} else if IsNumber(r[i]) {
			if twoDigit && !slash {
				return false
			}

			twoDigit = true

			continue
		}

		slash, twoDigit = false, false
	}

	return true
}

// RepeatRune is a function that repeats character count times
func RepeatRune(character rune, count int) []rune {
	result := make([]rune, 0, count)

	for i := 0; i < count; i++ {
		result = append(result, character)
	}

	return result
}

// UnpackString is a function that unpacks string in accordance with the rules specified in the task
func UnpackString(str string) (string, error) {
	if len(str) < 1 {
		return str, nil
	}

	runes := []rune(str)

	if !IsCorrect(runes) {
		return "", errCorrectness
	}

	unpackedString, slash := make([]rune, 0, len(runes)), false

	for i, value := range runes {
		if IsSlash(value) {
			if slash {
				unpackedString, slash = append(unpackedString, value), false

				continue
			}

			slash = true
		} else if IsNumber(value) {
			if slash {
				unpackedString, slash = append(unpackedString, value), false
			} else {
				repeatedRune := RepeatRune(runes[i-1], int(value-'0')-1)
				unpackedString = append(unpackedString, repeatedRune...)
			}
		} else {
			unpackedString = append(unpackedString, value)
		}
	}

	return string(unpackedString), nil
}

func main() {
	strs := []string{
		"a4bc2d5e",
		"abcd",
		"45",
		"",
		`qwe\4\5`,
		`qwe\45`,
		`qwe\\5`,
	}

	for i := range strs {
		unpacked, err := UnpackString(strs[i])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		fmt.Println(unpacked)
	}
}
