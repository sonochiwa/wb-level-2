package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func download(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка при запросе:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Ошибка при запросе. Код статуса:", resp.StatusCode)
		return
	}

	// Извлекаем имя файла из URL
	fileName := path.Base(url)

	// Создаем директорию для сохранения файлов
	err = os.MkdirAll(fileName, os.ModePerm)
	if err != nil {
		fmt.Println("Ошибка при создании директории:", err)
		return
	}

	// Создаем файл для сохранения данных
	filePath := path.Join(fileName, "index.html")
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer file.Close()

	// Копируем данные из ответа в файл
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("Ошибка при копировании данных:", err)
		return
	}

	fmt.Printf("Сайт успешно скачан в директорию: %s\n", fileName)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Использование: go run main.go <URL>")
		os.Exit(1)
	}

	url := os.Args[1]
	download(url)
}
