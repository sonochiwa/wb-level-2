package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var months = []string{
	"january",
	"february",
	"march",
	"april",
	"may",
	"june",
	"july",
	"august",
	"september",
	"november",
	"october",
	"december",
}

var si = map[string]int{
	"K": 3,
	"M": 6,
	"G": 9,
	"T": 12,
	"P": 15,
	"E": 18,
	"Z": 21,
	"Y": 24,
}

var (
	// Путь до файла, который мы будем сортировать
	filePath string

	// Поддерживаемые CLI флаги
	kFlag int
	nFlag bool
	rFlag bool
	uFlag bool
	mFlag bool
	bFlag bool
	cFlag bool
	hFlag bool
)

var rootCmd = &cobra.Command{
	Use:   "sort",
	Short: "Sort file",
	Long:  "This util sorts the file",
	Run: func(cmd *cobra.Command, args []string) {
		sortFile(args)
	},
}

// Функция сортировки файла
func sortFile(args []string) {
	// Если путь до файла не указан, то печатаем ошибку
	if len(args) == 0 {
		fmt.Println("file path not specified")
		return
	}

	filePath = args[0]
	data, err := readFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	if kFlag == 0 {
		fmt.Println("Invalid argument for -k")
		return
	}

	if kFlag > 0 {
		for i := 0; i < len(data); i++ {
			if kFlag > len(data[i]) {
				fmt.Println("Invalid argument for -k")

				printData(data)
				return
			}
		}
	}

	copyData := make([][]string, len(data), cap(data))
	if cFlag {
		copy(copyData, data)
	}

	if nFlag && mFlag {
		fmt.Println("Sort by numbers and months cannot be used together")
		return
	}

	if nFlag {
		sortNumbers(data)
	}

	if mFlag {
		data = sortMonths(data)
	}

	if hFlag {
		data = humanNumericSort(data)
	}

	if !nFlag && !mFlag && !hFlag {
		sortStrings(data)
	}

	if bFlag {
		data = trimSpace(data)
	}

	if rFlag {
		reverseData(data)
	}

	if uFlag {
		data = dedupeData(data)
	}

	if cFlag {
		compare(data, copyData)
	} else {
		printData(data)
	}
}

func humanNumericSort(data [][]string) [][]string {
	var numbers [][]string
	var shortening [][]string
	var result [][]string

	// Делим исходные данные на 2 массива - обычными числами и с сокращениями
	for i := 0; i < len(data); i++ {
		tail := data[i][kFlag-1][len(data[i][kFlag-1])-1:]
		_, exists := si[tail]
		if exists {
			shortening = append(shortening, data[i])
		} else {
			numbers = append(numbers, data[i])
		}
	}

	// Сортируем массив с обычными числами
	for i := 0; i < len(data)-1; i++ {
		for j := 0; j < len(data)-1; j++ {
			value1, _ := strconv.Atoi(data[j][kFlag-1])
			value2, _ := strconv.Atoi(data[j+1][kFlag-1])
			if value1 > value2 {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}

	// Сортируем сокращенные числа
	for i := 0; i < len(shortening)-1; i++ {
		for j := 0; j < len(shortening)-1; j++ {
			number1 := shortening[j][kFlag-1][:len(shortening[j][kFlag-1])-1]
			tail1 := shortening[j][kFlag-1][len(shortening[j][kFlag-1])-1:]
			tail1res := si[tail1]

			x1, _ := strconv.Atoi(number1)

			number2 := shortening[j+1][kFlag-1][:len(shortening[j+1][kFlag-1])-1]
			tail2 := shortening[j+1][kFlag-1][len(shortening[j+1][kFlag-1])-1:]
			tail2res := si[tail2]

			x2, _ := strconv.Atoi(number2)

			value1 := x1 * int(math.Pow(10, float64(tail1res)))
			value2 := x2 * int(math.Pow(10, float64(tail2res)))

			if value1 > value2 {
				shortening[j], shortening[j+1] = shortening[j+1], shortening[j]
			}
		}
	}

	// Мерджим слайсы
	for i := range numbers {
		result = append(result, numbers[i])
	}

	for i := range shortening {
		result = append(result, shortening[i])
	}

	return result
}

func compare(afterSortData [][]string, beforeSortData [][]string) {
	for i := 0; i < len(afterSortData); i++ {
		for j := 0; j < len(afterSortData[i]); j++ {
			if afterSortData[i][j] != beforeSortData[i][j] {
				fmt.Println("disorder:", strings.Join(afterSortData[i], " "))
				return
			}
		}
	}
}

func readFile(filePath string) ([][]string, error) {
	var data [][]string

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Кастомный сплит строки, чтобы оставались висячие пробелы
		parsedData := splitString(scanner.Text())
		data = append(data, parsedData)
	}
	return data, nil
}

func splitString(input string) []string {
	var result []string
	var currentWord string

	for i, char := range input {
		if char != ' ' {
			currentWord += string(char)
		} else {
			if input[i-1] == ' ' {
				currentWord += string(char)
			} else {
				result = append(result, currentWord)
				currentWord = ""
			}
		}
	}

	// Добавляем последнее слово, если оно есть
	if currentWord != "" {
		result = append(result, currentWord)
	}

	return result
}

// TODO: заменить на quickSort
func sortStrings(data [][]string) {
	for i := 0; i < len(data)-1; i++ {
		for j := 0; j < len(data)-1; j++ {
			if data[j][kFlag-1] > data[j+1][kFlag-1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}

func sortNumbers(data [][]string) {
	for i := 0; i < len(data)-1; i++ {
		for j := 0; j < len(data)-1; j++ {
			value1, _ := strconv.Atoi(data[j][kFlag-1])
			value2, _ := strconv.Atoi(data[j+1][kFlag-1])
			if value1 > value2 {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}

func sortMonths(data [][]string) [][]string {
	var result [][]string

	for i := 0; i < len(months); i++ {
		for j := 0; j < len(data); j++ {
			if strings.ToLower(strings.TrimSpace(data[j][kFlag-1])) == months[i] {
				result = append(result, data[j])
			}
		}
	}

	return result
}

func reverseData(data [][]string) {
	for i := 0; i < len(data)/2; i++ {
		data[i], data[len(data)-i-1] = data[len(data)-i-1], data[i]
	}
}

func trimSpace(data [][]string) [][]string {
	var result [][]string
	for i := range data {
		var str []string
		for j := 0; j < len(data[i]); j++ {
			str = append(str, strings.TrimSpace(data[i][j]))
		}
		result = append(result, str)
	}

	return result
}

func dedupeData(data [][]string) [][]string {
	dataMap := make(map[string]bool)
	var uniqueValues []string
	var result [][]string

	for i := range data {
		value := strings.Join(data[i], " ")
		_, exists := dataMap[value]
		if !exists {
			uniqueValues = append(uniqueValues, value)
			dataMap[value] = true
		}
	}

	for i := range uniqueValues {
		result = append(result, strings.SplitAfter(uniqueValues[i], ", "))
	}

	return result
}

func printData(data [][]string) {
	for _, v := range data {
		fmt.Printf("%s\n", strings.Join(v, " "))
	}
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&kFlag, "key", "k", 1, "")
	rootCmd.PersistentFlags().BoolVarP(&nFlag, "numeric-sort", "n", false, "")
	rootCmd.PersistentFlags().BoolVarP(&rFlag, "reverse", "r", false, "")
	rootCmd.PersistentFlags().BoolVarP(&uFlag, "unique", "u", false, "")
	rootCmd.PersistentFlags().BoolVarP(&mFlag, "month-sort", "M", false, "")
	rootCmd.PersistentFlags().BoolVarP(&bFlag, "ignore-leading-blanks", "b", false, "")
	rootCmd.PersistentFlags().BoolVarP(&cFlag, "check", "c", false, "")
	rootCmd.PersistentFlags().BoolVarP(&hFlag, "human-numeric-sort", "H", false, "")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
