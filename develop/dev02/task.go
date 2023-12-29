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

// Метод unpack распаковывает строку, содержащую повторяющиеся символы
func unpack(s []rune) string {
	// Планируется частая конкатенация, поэтому используем string builder
	var result strings.Builder

	// Итерируем по массиву рун, для поддержки unicode символов
	for i := 0; i < len(s); i++ {
		// Если s[i] не экранирующий символ
		if s[i] == '\\' {
			result.WriteString(string(s[i+1]))
			i++
		} else {
			// Если s[i] число, то
			if unicode.IsDigit(s[i]) {
				// Конвертируем в int
				count, _ := strconv.Atoi(string(s[i]))
				// Распаковываем и записываем в result
				result.WriteString(strings.Repeat(string(s[i-1]), count-1))
			} else {
				// Если s[i] не число, тогда пишем его в result
				result.WriteString(string(s[i]))
			}
		}
	}

	// Возвращаем распакованную строку
	return result.String()
}

func main() {
	// Входные данные
	s := []rune("qwe\\4\\5")

	// Распаковываем строку
	result := unpack(s)

	// Выводим результат
	fmt.Println(result)
}
