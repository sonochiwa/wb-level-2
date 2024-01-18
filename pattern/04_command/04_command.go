package main

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

/*
Команда (Command) - паттерн поведения объектов

Инкапсулирует запрос в объекте, позволяя тем самым параметризировать клиенты для разных запросов,
ставить запросы в очередь или протоколировать их, а также поддерживать отмену операций.
*/

// Command - интерфейс команды
type Command interface {
	Execute()
}

// ConcreteCommand - реализация конкретной команды
type ConcreteCommand struct {
	receiver Receiver
}

func (c *ConcreteCommand) Execute() {
	c.receiver.Action()
}

// Receiver - получатель
type Receiver struct{}

func (r *Receiver) Action() {
	fmt.Println("Выполняется действие")
}

// Invoker - инвокер (Исполнитель команд)
type Invoker struct {
	command Command
}

func (i *Invoker) SetCommand(command Command) {
	i.command = command
}

func (i *Invoker) ExecuteCommand() {
	i.command.Execute()
}

func NewReceiver() Receiver {
	return Receiver{}
}

func main() {
	// Создание объектов
	receiver := NewReceiver()
	command := &ConcreteCommand{receiver: receiver}
	invoker := Invoker{}

	// Назначение команды исполнителю
	invoker.SetCommand(command)

	// Выполнение команды
	invoker.ExecuteCommand()
}
