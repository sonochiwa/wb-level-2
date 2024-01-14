package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Поддерживаемые CLI флаги
var (
	afterFlag      int
	beforeFlag     int
	contextFlag    int
	countFlag      bool
	ignoreCaseFlag bool
	invertFlag     bool
	lineNumFlag    bool
)

var rootCmd = &cobra.Command{
	Use:   "grep",
	Short: "Grep data from file",
	Long:  "This util grep data from file",
	Run: func(cmd *cobra.Command, args []string) {
		grepData(args)
	},
}

func grepData(args []string) {
	var filePath string // Путь до файла, который мы будем сортировать
	var pattern string  // Шаблон для поиска
	var indexes map[int]bool

	// Если не указан шаблон или путь до файла, то печатаем ошибку
	if len(args) == 2 {
		pattern = args[0]
		filePath = args[1]
	} else {
		fmt.Println("file path or pattern not specified")
		return
	}

	data, err := parseData(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	indexes = getIndexes(pattern, data)

	if afterFlag > 0 {
		after(indexes, data, pattern, afterFlag)
	}

	if beforeFlag > 0 {
		before(indexes, data, pattern, beforeFlag)
	}

	if contextFlag > 0 {
		context(indexes, data, pattern, contextFlag)
	}

	if invertFlag {
		invert(indexes)
	}

	if countFlag {
		count := countStrings(indexes)
		fmt.Println(count)
		return
	}

	printData(data, indexes)
}

// Функция для чтения строк из файла
func parseData(filePath string) ([]string, error) {
	var data []string

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data, nil
}

// Функция для заполнения мапы индексов
func getIndexes(pattern string, data []string) map[int]bool {
	indexes := make(map[int]bool, len(data))

	for i := 0; i < len(data); i++ {
		normalizedData := data[i]
		normalizePattern := pattern

		if ignoreCaseFlag {
			normalizedData = strings.ToLower(normalizedData)
			normalizePattern = strings.ToLower(pattern)
		}

		if strings.Contains(normalizedData, normalizePattern) {
			indexes[i] = true
		} else {
			indexes[i] = false
		}
	}

	return indexes
}

// Функция для обработки флага -A (after)
func after(indexes map[int]bool, data []string, pattern string, offset int) {
	var keys []int
	var maxKey int

	for i := range indexes {
		if indexes[i] == true && strings.Contains(data[i], pattern) {
			keys = append(keys, i)
		}
	}

	maxKey = slices.Max(keys)

	for i := 0; i < offset; i++ {
		if _, exists := indexes[maxKey+i+1]; exists {
			indexes[maxKey+i+1] = true
		}
	}
}

// Функция для обработки флага -B (before)
func before(indexes map[int]bool, data []string, pattern string, offset int) {
	var keys []int
	var minKey int

	for i := range indexes {
		if indexes[i] == true && strings.Contains(data[i], pattern) {
			keys = append(keys, i)
		}
	}

	minKey = slices.Min(keys)

	for i := 0; i < offset; i++ {
		if _, exists := indexes[minKey-i-1]; exists {
			indexes[minKey-i-1] = true
		}
	}
}

// Функция для обработки флага -C (context)
func context(indexes map[int]bool, data []string, pattern string, offset int) {
	var keys []int
	var minKey int
	var maxKey int

	for i := range indexes {
		if indexes[i] == true && strings.Contains(data[i], pattern) {
			keys = append(keys, i)
		}
	}

	minKey = slices.Min(keys)
	maxKey = slices.Max(keys)

	for i := 0; i < offset; i++ {
		if _, exists := indexes[minKey-i-1]; exists {
			indexes[minKey-i-1] = true
		}
		if _, exists := indexes[maxKey+i+1]; exists {
			indexes[maxKey+i+1] = true
		}
	}
}

// Функция для инвертирования значений индексов в мапе
func invert(indexes map[int]bool) {
	for i := range indexes {
		if indexes[i] == true {
			indexes[i] = false
		} else {
			indexes[i] = true
		}
	}
}

// Функция для подсчета количества строк, которые нужно вывести
func countStrings(indexes map[int]bool) int {
	var count int

	for i := range indexes {
		if indexes[i] == true {
			count++
		}
	}

	return count
}

// Функция для печати данных, которые нужно вывести
func printData(data []string, indexes map[int]bool) {
	for i := range data {
		if key := indexes[i]; key == true {
			if lineNumFlag {
				fmt.Println(strconv.Itoa(i+1) + ":" + data[i])
			} else {
				fmt.Println(data[i])
			}
		}
	}
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&afterFlag, "1", "A", 0, "")
	rootCmd.PersistentFlags().IntVarP(&beforeFlag, "2", "B", 0, "")
	rootCmd.PersistentFlags().IntVarP(&contextFlag, "3", "C", 0, "")
	rootCmd.PersistentFlags().BoolVarP(&countFlag, "4", "c", false, "")
	rootCmd.PersistentFlags().BoolVarP(&ignoreCaseFlag, "5", "i", false, "")
	rootCmd.PersistentFlags().BoolVarP(&invertFlag, "6", "v", false, "")
	rootCmd.PersistentFlags().BoolVarP(&lineNumFlag, "7", "n", false, "")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
