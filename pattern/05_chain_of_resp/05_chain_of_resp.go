package main

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

/*
Цепочка обязанностей (Chain of Responsibility) - паттерн поведения объектов

Позволяет избежать привязки отправителя запроса к его получателю, предоставляя возможность
обработать запрос несколькими объектами. Связывает объекты-получатели в цепочку и передает
запрос по этой цепочке, пока он не будет обработан

ПРИМЕНИМОСТЬ

- Запрос может быть обработан более чем одним объектом, причем настоящий обработчик заранее
неизвестен и должен быть найден автоматически

- Набор объектов, способных обрабатывать запрос, должен задаваться динамически

УЧАСТНИКИ

- Handler - обработчик
-- определяет интерфейс для обработки запросов
-- реализует связь с преемником

- ConcreteHandler - конкретный обработчик
-- обрабатывает запрос, за который отвечает
-- имеет доступ к своему преемнику
-- если ConcreteHandler способен обрабатывать запрос, то так и делает,
если не может, то направляет его своему преемнику

- Client - клиент
-- отправляет запрос некоторому объекту ConcreteHandler в цепочке

РЕЗУЛЬТАТ

- ослабление связанности. Паттерн освобождает объект от необходимости знать,
кто конкретно обрабатывает его запрос. Отправитель и получатель ничего не знают
друг о друге, а включенный в цепочку объект - о структуре цепочки

В результате цепочка событий помогает упростить взаимосвязи между объектами

- дополнительная гибкость при распределении обязанностей между объектами.
Паттерн позволяет повысить гибкость распределения обязанностей между объектами.
Добавить или изменить обязанности по обработке запроса можно, включив в цепочку
новых участников или изменив ее каким-то другим образом.

- получение не гарантировано. Поскольку у запроса нет явного получателя, то нет
и гарантий, что он вообще будет обработан: он может достичь конца цепочки и пропасть.

*/

// Handler - интерфейс обработчика
type Handler interface {
	HandleRequest(request int)
	SetNextHandler(handler Handler)
}

// ConcreteHandlerA - конкретный обработчик A
type ConcreteHandlerA struct {
	nextHandler Handler
}

func (c *ConcreteHandlerA) HandleRequest(request int) {
	if request <= 10 {
		fmt.Println("ConcreteHandlerA обрабатывает запрос", request)
	} else if c.nextHandler != nil {
		c.nextHandler.HandleRequest(request)
	}
}

func (c *ConcreteHandlerA) SetNextHandler(handler Handler) {
	c.nextHandler = handler
}

// ConcreteHandlerB - конкретный обработчик B
type ConcreteHandlerB struct {
	nextHandler Handler
}

func (c *ConcreteHandlerB) HandleRequest(request int) {
	if request > 10 && request <= 20 {
		fmt.Println("ConcreteHandlerB обрабатывает запрос", request)
	} else if c.nextHandler != nil {
		c.nextHandler.HandleRequest(request)
	}
}

func (c *ConcreteHandlerB) SetNextHandler(handler Handler) {
	c.nextHandler = handler
}

// ClientCode - клиентский код
func ClientCode(handler Handler, requests []int) {
	for _, request := range requests {
		handler.HandleRequest(request)
	}
}

func main() {
	handlerA := &ConcreteHandlerA{}
	handlerB := &ConcreteHandlerB{}

	// Формирование цепочки обработчиков
	handlerA.SetNextHandler(handlerB)

	// Обработка запросов
	requests := []int{5, 15, 25}
	ClientCode(handlerA, requests)
}
