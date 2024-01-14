package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestDownload(t *testing.T) {
	// Создаем временный HTTP сервер для тестов
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, Test!"))
	}))
	defer ts.Close()

	// Получаем адрес временного сервера
	url := ts.URL

	// Вызываем функцию download
	download(url)

	// Проверяем, что директория была создана
	dirName := filepath.Base(url)
	_, err := os.Stat(dirName)
	if os.IsNotExist(err) {
		t.Errorf("Ожидалось, что директория %s будет создана, но не создана", dirName)
	}

	// Проверяем, что файл был создан и содержит ожидаемые данные
	filePath := filepath.Join(dirName, "index.html")
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		t.Errorf("Ошибка при чтении файла %s: %v", filePath, err)
	}

	expectedContent := "Hello, Test!"
	if string(fileContent) != expectedContent {
		t.Errorf("Ожидалось, что содержимое файла %s будет '%s', но получено '%s'", filePath, expectedContent, string(fileContent))
	}
}
