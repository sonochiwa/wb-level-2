package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("До свидания!")
				os.Exit(0)
			}
			fmt.Fprintln(os.Stderr, "Ошибка чтения ввода:", err)
			continue
		}

		// Удаление символа новой строки из ввода
		input = strings.TrimSuffix(input, "\n")

		// Разбивка введенной строки на команды
		commands := strings.Fields(input)

		if len(commands) == 0 {
			continue
		}

		switch commands[0] {
		case "cd":
			if len(commands) < 2 {
				fmt.Println("Используйте: cd <директория>")
			} else {
				err := os.Chdir(commands[1])
				if err != nil {
					fmt.Fprintln(os.Stderr, "Ошибка при смене директории:", err)
				}
			}
		case "pwd":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Ошибка при получении текущей директории:", err)
			} else {
				fmt.Println(dir)
			}
		case "echo":
			fmt.Println(strings.Join(commands[1:], " "))
		case "kill":
			if len(commands) < 2 {
				fmt.Println("Используйте: kill <PID>")
			} else {
				pid := commands[1]
				err := exec.Command("kill", pid).Run()
				if err != nil {
					fmt.Fprintln(os.Stderr, "Ошибка при отправке сигнала:", err)
				}
			}

		case "ps":
			cmd := exec.Command("ps", "aux")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Ошибка при выполнении команды ps:", err)
			}
		case "fork-exec":
			if len(commands) < 2 {
				fmt.Println("Используйте: fork-exec <команда>")
			} else {
				cmd := exec.Command(commands[1], commands[2:]...)
				err := cmd.Start()
				if err != nil {
					fmt.Fprintln(os.Stderr, "Ошибка при запуске процесса:", err)
				}
			}

		default:
			cmd := exec.Command(commands[0], commands[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Ошибка при выполнении команды:", err)
			}
		}
	}
}
