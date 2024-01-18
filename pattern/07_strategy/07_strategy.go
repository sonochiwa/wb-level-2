package main

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

/*
Стратегия (Strategy) - поведенческий паттерн

Представляет собой поведенческий шаблон проектирования, который определяет семейство алгоритмов,
инкапсулирует каждый из них и делает их взаимозаменяемыми. Паттерн стратегия позволяет выбирать
алгоритм в зависимости от контекста
*/

// PaymentStrategy - интерфейс стратегии
type PaymentStrategy interface {
	Pay(amount float64) string
}

// CreditCardPayment - конкретная стратегия 1
type CreditCardPayment struct{}

func (c *CreditCardPayment) Pay(amount float64) string {
	return fmt.Sprintf("Оплата %.2f с использованием кредитной карты", amount)
}

// Pay - конкретная стратегия 2
type PayPalPayment struct{}

func (p *PayPalPayment) Pay(amount float64) string {
	return fmt.Sprintf("Оплата %.2f с использованием PayPal", amount)
}

// ShoppingCart - контекст
type ShoppingCart struct {
	PaymentStrategy PaymentStrategy
}

// SetPaymentStrategy - метод для установки стратегии оплаты
func (s *ShoppingCart) SetPaymentStrategy(paymentStrategy PaymentStrategy) {
	s.PaymentStrategy = paymentStrategy
}

// ProcessPayment - метод для выполнения оплаты
func (s *ShoppingCart) ProcessPayment(amount float64) string {
	if s.PaymentStrategy == nil {
		return "Выберите метод оплаты"
	}
	return s.PaymentStrategy.Pay(amount)
}

func main() {
	// Пример использования

	// Создание объекта корзины
	cart := ShoppingCart{}

	// Установка стратегии оплаты (кредитная карта)
	cart.SetPaymentStrategy(&CreditCardPayment{})

	// Выполнение оплаты
	result := cart.ProcessPayment(100.50)
	fmt.Println(result)

	// Изменение стратегии оплаты (PayPal)
	cart.SetPaymentStrategy(&PayPalPayment{})

	// Выполнение оплаты с новой стратегией
	result = cart.ProcessPayment(75.25)
	fmt.Println(result)
}
