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
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Поддерживаемые CLI флаги
var (
	fieldsFlag    string
	delimiterFlag string
	separatedFlag bool
)

var rootCmd = &cobra.Command{
	Use:   "cut",
	Short: "Cut data from file",
	Long:  "This util cuts data from file",
	Run: func(cmd *cobra.Command, args []string) {
		cutData(args)
	},
}

func cutData(args []string) {
	var filePath string // Путь до файла, который мы будем сортировать

	// Если не указан шаблон или путь до файла, то печатаем ошибку
	if len(args) == 1 {
		filePath = args[0]
	} else {
		fmt.Println("file path not specified")
		return
	}

	data, err := parseData(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(fieldsFlag) < 1 {
		fmt.Println("Не указаны колонки для вывода")
		return
	}

	fields := parseFields()

	// Если значение колонки больше количества колонок в первой строке, то печатаем ошибку
	if slices.Max(fields) > len(data[0]) {
		fmt.Println("Некорректное значение для флага -f")
		return
	}

	result := selectedFields(data, fields)

	printData(result)
}

// Функция для чтения строк из файла
func parseData(filePath string) ([][]string, error) {
	var data [][]string

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if separatedFlag && !strings.Contains(scanner.Text(), delimiterFlag) {
			continue
		}

		parsedData := strings.Split(scanner.Text(), delimiterFlag)
		data = append(data, parsedData)
	}
	return data, nil
}

func printData(data [][]string) {
	for _, v := range data {
		fmt.Printf("%s\n", strings.Join(v, delimiterFlag))
	}
}

func parseFields() []int {
	var fields []int

	strFields := strings.Split(fieldsFlag, ",")

	for _, v := range strFields {
		key, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		fields = append(fields, key)
	}

	return fields
}

func selectedFields(data [][]string, fields []int) [][]string {
	var result [][]string

	for i := range data {
		var currentRow []string
		for j := range fields {
			currentRow = append(currentRow, data[i][fields[j]-1])
		}
		result = append(result, currentRow)
	}

	return result
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&fieldsFlag, "1", "f", "", "")
	rootCmd.PersistentFlags().StringVarP(&delimiterFlag, "2", "d", "	", "")
	rootCmd.PersistentFlags().BoolVarP(&separatedFlag, "3", "s", false, "")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
