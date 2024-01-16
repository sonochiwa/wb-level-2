package main

import (
	"errors"
	"fmt"
	"time"
)

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

/*
Фасад (Facade) - Структурный паттернр проектирования. Представляет собой
простой доступ (просто интерфейс) к сложнеой системе.

Когда у нас есть много разных подсистем котоыре используют свои интерфейсы
и реализуют какой-то свой функционал поведения.

Преимущества:
- Изолирует клиентов от поведения сложной системы
- Сам интерфейс фасада простой

Минусы:
- Является супер-объектом и все последующие вызовы в системе будут проходить
через этот объект
*/

type Product struct {
	Name  string
	Price float64
}

type Shop struct {
	Name     string
	Products []Product
}

type Bank struct {
	Name  string
	Cards []Card
}

type Card struct {
	Name    string
	Balance float64
	Bank    *Bank
}

type User struct {
	Name string
	Card *Card
}

func (user User) GetBalance() float64 {
	return user.Card.Balance
}

func (shop Shop) Sell(user User, product string) error {
	println("[Магазин] Запрос к пользователю, для получения отстатка по карте")
	time.Sleep(500 * time.Millisecond)
	err := user.Card.CheckBalance()
	if err != nil {
		return err
	}

	fmt.Printf("[Магазин] Проверка - может ли [%s], купить товар\n", user.Name)
	time.Sleep(500 * time.Millisecond)
	for _, prod := range shop.Products {
		if prod.Name != product {
			continue
		}

		if prod.Price > user.GetBalance() {
			return errors.New("[Магазин] Недостаточно средств для покупки товара")
		}

		fmt.Printf("[Магазин] Товар [%s] - куплен!\n", prod.Name)
	}

	return nil
}

func (card Card) CheckBalance() error {
	println("[Карта] Запрос в банк для проверки остатка")
	time.Sleep(800 * time.Millisecond)

	return card.Bank.CheckBalance(card.Name)
}

func (bank Bank) CheckBalance(cardNumber string) error {
	println(fmt.Sprintf("[Банк] Получение остатка по карте %s", cardNumber))
	time.Sleep(300 * time.Millisecond)

	for _, card := range bank.Cards {
		if card.Name != cardNumber {
			continue
		}
		if card.Balance <= 0 {
			return errors.New("[Банк] Недостаточно средств")
		}
	}

	println("[Банк] Остаток положительный!")

	return nil
}

var (
	bank = Bank{
		Name:  "БАНК",
		Cards: []Card{},
	}

	card1 = Card{
		Name:    "CRD-1",
		Balance: 200,
		Bank:    &bank,
	}

	card2 = Card{
		Name:    "CRD-2",
		Balance: 5,
		Bank:    &bank,
	}

	user1 = User{
		Name: "Покупатель-1",
		Card: &card1,
	}

	user2 = User{
		Name: "Покупатель-2",
		Card: &card2,
	}

	prod = Product{
		Name:  "Сыр",
		Price: 150,
	}

	shop = Shop{
		Name: "SHOP",
		Products: []Product{
			prod,
		},
	}
)

func main() {
	println("[Банк] Выпуск карт")
	bank.Cards = append(bank.Cards, card1, card2)

	fmt.Printf("[%s]\n", user1.Name)

	err := shop.Sell(user1, prod.Name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("[%s]\n", user2.Name)
	err = shop.Sell(user2, prod.Name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
