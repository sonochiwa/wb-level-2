package main

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
Строитель (Builder) - порождающий объекты паттерн.

Специализируется на пошаговом построении объекта

Реализация может выглядеть как пайплайн методов для создания сложного объекта.
Если описать кратко, то билдер это пайплайн из сеттеров где последний шаг создает объект ( .Build() )

УЧАСТНИКИ

Builder - строитель:
- Задает абстрактный интерфейс для создания частей объекта Product

ConcreteBuilder - конкретный строитель:
- конструирует и собирает вместе части продукта посредством реализации интерфейса Builder
- предоставляет интерфейс для доступа к продукту

Director - распорядитель:
- конструирует объект, пользуясь интерфейсом Builder

Product - продукт:
- представляет сложный конструируемый объект. ConcreteBuilder строит внутреннее
представление продукта и определяет процесс его сборки
- включает классы, которые определяют составные части, в том числе интерфейсы для сборки
конечного результата из частей

ОТНОШЕНИЯ

- клиент создает объект-распорядитель Director и настраивает его нужным объектом-строителем Builder
- распорядитель уведомляет строителя о том, что нужно построить очередную часть продукта
- клиент забирает продукт у строителя

ПРЕИМУЩЕСТВА

- паттерн позволяет изменять внутреннее поведение продукта
- изолирует код, реализующий конструирование и представление
- предоставляет более точный контроль над процессом конструирование

Похожим паттерном является абстрактная фабрика, но она только она специализируется на создании
семейств объектов (простых или сложных)

Строитель возвращает продукт на последнем шаге, а абстрактная фабрика немедленно
*/

type ComputerBuilderI interface {
	CPU(val string) ComputerBuilderI
	RAM(val int) ComputerBuilderI
	MB(val string) ComputerBuilderI

	Build() Computer
}

type Computer struct {
	CPU string
	RAM int
	MB  string
}

type computerBuilder struct {
	cpu string
	ram int
	mb  string
}

func (b *computerBuilder) CPU(val string) ComputerBuilderI {
	b.cpu = val
	return b
}

func (b *computerBuilder) RAM(val int) ComputerBuilderI {
	b.ram = val
	return b
}

func (b *computerBuilder) MB(val string) ComputerBuilderI {
	b.mb = val
	return b
}

func (b computerBuilder) Build() Computer {
	return Computer{
		CPU: b.cpu,
		RAM: b.ram,
		MB:  b.mb,
	}
}

func NewComputerBuilder() ComputerBuilderI {
	return &computerBuilder{}
}

func main() {
	compBuilder := NewComputerBuilder()
	computer := compBuilder.CPU("Ryzen 5600x").RAM(16).MB("GigaByte").Build()
	fmt.Println(computer)
}
