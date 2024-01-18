package main

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

/*
Фабричный метод (Factory Method) - порождающий объекты паттерн

Представляет собой интерфейс для создания экземпляра некоторого класса, но оставляет
подклассам решение о том, какой класс инстанцировать
*/

// Product - интерфейс для создания продуктов
type Product interface {
	GetName() string
}

// ConcreteProduct1 - конкретный продукт 1
type ConcreteProduct1 struct{}

func (p *ConcreteProduct1) GetName() string {
	return "ConcreteProduct1"
}

// ConcreteProduct2 - конкретный продукт 2
type ConcreteProduct2 struct{}

func (p *ConcreteProduct2) GetName() string {
	return "ConcreteProduct2"
}

// Factory - интерфейс для создания фабрики
type Factory interface {
	CreateProduct() Product
}

// ConcreteFactory1 - конкретная фабрика 1
type ConcreteFactory1 struct{}

func (f *ConcreteFactory1) CreateProduct() Product {
	return &ConcreteProduct1{}
}

// ConcreteFactory2 - конкретная фабрика 2
type ConcreteFactory2 struct{}

func (f *ConcreteFactory2) CreateProduct() Product {
	return &ConcreteProduct2{}
}

func main() {
	// Используем фабрику 1
	factory1 := &ConcreteFactory1{}
	product1 := factory1.CreateProduct()
	fmt.Println(product1.GetName())

	// Используем фабрику 2
	factory2 := &ConcreteFactory2{}
	product2 := factory2.CreateProduct()
	fmt.Println(product2.GetName())
}
