package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func unpack(s string) (string, error) {
	var result strings.Builder

	for i, char := range s {
		if unicode.IsDigit(char) {
			count, _ := strconv.Atoi(string(char))
			result.WriteString(strings.Repeat(string(s[i-1]), count-1))
		} else {
			result.WriteString(string(char))
		}
	}

	return result.String(), nil
}

func main() {
	s := "a4bc2d5e"
	result, _ := unpack(s)
	fmt.Println(result)
}
