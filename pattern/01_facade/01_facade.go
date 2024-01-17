package main

import (
	"fmt"
)

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

/*
Фасад (Facade) - структурный паттерн.
Представляет собой простой доступ (простой интерфейс) к сложной системе.

Используем его, когда у нас есть много разных подсистем, которые используют свои интерфейсы
и реализуют какой-то свой функционал поведения.

Реализует принципы SOLID:
- Открытости/закрытости (интерфейс в который скрывает код подсистемы)
- Инверсия зависимостей (уменьшает связанность)

Преимущества:
- Изолирует клиентов от поведения сложной системы
- Сам интерфейс фасада простой

Минусы:
- Является супер-объектом и все последующие вызовы в системе будут проходить
через этот объект
*/

// SubsystemA - подсистема A
type SubsystemA struct {
}

func (s *SubsystemA) OperationA() {
	fmt.Println("Subsystem A: Operation A")
}

// SubsystemB - подсистема B
type SubsystemB struct {
}

func (s *SubsystemB) OperationB() {
	fmt.Println("Subsystem B: Operation B")
}

// SubsystemC - подсистема C
type SubsystemC struct {
}

func (s *SubsystemC) OperationC() {
	fmt.Println("Subsystem C: Operation C")
}

// Facade - фасад
type Facade struct {
	subsystemA *SubsystemA
	subsystemB *SubsystemB
	subsystemC *SubsystemC
}

func NewFacade() *Facade {
	return &Facade{
		subsystemA: &SubsystemA{},
		subsystemB: &SubsystemB{},
		subsystemC: &SubsystemC{},
	}
}

// FacadeOperation - высокоуровневая операция, предоставляемая фасадом
func (f *Facade) FacadeOperation() {
	fmt.Println("Facade Operation:")
	f.subsystemA.OperationA()
	f.subsystemB.OperationB()
	f.subsystemC.OperationC()
}

func main() {
	// Использование фасада для упрощения взаимодействия с подсистемой
	facade := NewFacade()
	facade.FacadeOperation()
}
