package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {
	// Парсим аргументы командной строки
	host := flag.String("host", "", "Хост для подключения")
	port := flag.Int("port", 0, "Порт для подключения")
	timeout := flag.Duration("timeout", 10*time.Second, "Таймаут подключения")
	flag.Parse()

	// Проверяем обязательные аргументы
	if *host == "" || *port == 0 {
		fmt.Println("Необходимо указать хост и порт")
		return
	}

	// Формируем адрес для подключения
	address := fmt.Sprintf("%s:%d", *host, *port)

	// Устанавливаем соединение
	conn, err := net.DialTimeout("tcp", address, *timeout)
	if err != nil {
		fmt.Printf("Ошибка подключения: %v\n", err)
		return
	}
	defer conn.Close()

	// Запускаем горутину для чтения данных из сокета и вывода в STDOUT
	go func() {
		io.Copy(os.Stdout, conn)
		fmt.Println("Сервер закрыл соединение")
		os.Exit(0)
	}()

	// Запускаем горутину для чтения данных из STDIN и записи в сокет
	go func() {
		io.Copy(conn, os.Stdin)
		fmt.Println("Выход по Ctrl+D")
		os.Exit(0)
	}()

	// Ждем сигнала завершения программы (Ctrl+C)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	fmt.Println("Программа завершена")
}
