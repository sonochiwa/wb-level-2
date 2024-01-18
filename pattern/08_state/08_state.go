package main

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

/*
Стратегия (Strategy) - поведенческий паттерн

Может быть реализован для управления состоянием объекта.
Позволяет объекту изменять свое поведение при изменении его
внутреннего состояния
*/

// State - интерфейс состояния
type State interface {
	Handle() string
}

// StateA - реализации состояний
type StateA struct{}

func (s *StateA) Handle() string {
	return "State A"
}

type StateB struct{}

func (s *StateB) Handle() string {
	return "State B"
}

// Context - контекст, который использует состояния
type Context struct {
	state State
}

func (c *Context) SetState(state State) {
	c.state = state
}

func (c *Context) Request() string {
	return c.state.Handle()
}

func main() {
	context := &Context{}

	stateA := &StateA{}
	context.SetState(stateA)
	fmt.Println(context.Request()) // Вывод: State A

	stateB := &StateB{}
	context.SetState(stateB)
	fmt.Println(context.Request()) // Вывод: State B
}
