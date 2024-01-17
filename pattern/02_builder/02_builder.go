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
Если описать кратко, то билдер это пайплайн из сеттеров, где последний шаг создает объект ( .Build() )

УЧАСТНИКИ

Builder - строитель:
- Задает абстрактный интерфейс для создания частей объекта Product

ConcreteBuilder - конкретный строитель:
- конструирует и собирает вместе части продукта посредством реализации интерфейса Builder
- предоставляет интерфейс для доступа к продукту (в контексте го - структура)

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

// Builder - интерфейс строителя, определяющий шаги создания продукта
type Builder interface {
	BuildName(value string) Builder
	BuildPrice(value int) Builder

	GetProduct() Product
}

// Product - объект который мы строим
type Product struct {
	Name  string
	Price int
}

// ConcreteBuilder - конкретный строитель
type ConcreteBuilder struct {
	product Product
}

// Director - управляет процессом конструирования
type Director struct {
	builder Builder
}

func (b *ConcreteBuilder) BuildName(value string) Builder {
	b.product.Name = value
	return b
}

func (b *ConcreteBuilder) BuildPrice(value int) Builder {
	b.product.Price = value
	return b
}

func (b *ConcreteBuilder) GetProduct() Product {
	return b.product
}

func (d *Director) Construct() Product {
	d.builder.BuildName("Lemon")
	d.builder.BuildPrice(20)

	return d.builder.GetProduct()
}

// NewConcreteBuilder - конструктор конкретного билдера
func NewConcreteBuilder() Builder {
	return &ConcreteBuilder{}
}

// NewDirector - конструктор директора
func NewDirector(builder Builder) *Director {
	return &Director{builder: builder}
}

func main() {
	// Создаем ConcreteBuilder
	builder := NewConcreteBuilder()

	// Создаем директора и передаем ему строителя
	director := NewDirector(builder)

	// Директор управляет процессом конструирования
	product := director.Construct()

	fmt.Printf("%s, %d\n", product.Name, product.Price)
}
