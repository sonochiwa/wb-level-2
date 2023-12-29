package main

import (
	"errors"
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
func unpack(data string) (string, error) {
	s := []rune(data)

	// Планируется частая конкатенация, поэтому используем string builder
	var result strings.Builder

	// Итерируем по массиву рун, для поддержки unicode символов
	for i := 0; i < len(s); i++ {

		if s[len(s)-2] != '\\' && s[len(s)-1] == '\\' {
			return "", errors.New("некорректная строка")
		}

		if unicode.IsDigit(s[i]) && unicode.IsDigit(s[i+1]) {
			return "", errors.New("некорректная строка")
		}

		if unicode.IsDigit(s[0]) {
			return "", errors.New("некорректная строка")
		}

		// Если s[i] escape последовательность
		if s[i] == '\\' {
			result.WriteString(string(s[i+1]))
			// Инкрементируем счетчик i, чтобы пропустить следующий шаг
			i++
		} else {
			// Если s[i] число, то
			if unicode.IsDigit(s[i]) {
				// Конвертируем в int
				count, _ := strconv.Atoi(string(s[i]))
				// Распаковываем и записываем в result
				result.WriteString(strings.Repeat(string(s[i-1]), count-1))
			} else {
				// Если s[i] не число и не escape последовательность, тогда пишем его в result
				result.WriteString(string(s[i]))
			}
		}
	}

	// Возвращаем распакованную строку
	return result.String(), nil
}

func main() {
	// Входные данные
	s := "qwe\\\\5"

	// Распаковка строки
	result, err := unpack(s)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Вывод результата
	fmt.Println(result)
}
