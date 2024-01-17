package main

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

/*
Посетитель (visitor) - паттерн поведения объектов (поведенческий).

Позволяет определить новую операцию, не изменяя классы элементов, на которых эта
операция выполняется.

ПРИМЕНИМОСТЬ
- в структуре присутствуют объекты многих классов с различными интерфейсами, и вы хотите
выполнять над ними операции, зависящие от конкретных классов

- над объектами входящими в состав структуры должны выполняться разнообразные, не
связанные между собой операции и вы не хотите "засорять" классы. Посетитель позволяет
объединить родственные операции, поместив их в один класс.

УЧАСТНИКИ

- Visitor (NodeVisitor) - посетитель
Объявляет операцию Visit для каждого класса ConcreteElement в структуре объектов.
Имя и сигнатура этой операции позволяют пользователю определить, элемент какого конкретного
класса он посещает. Владея такой информацией, посетитель может обращаться к элементу напрямую
через его интерфейс

- ConcreteVisitor (TypeCheckingVisitor) - конкретный посетитель
Реализует все операции, объявленные в классе Visitor. Каждая операция реализует фрагмент
алгоритма, определенного для класса соответствующего объекта в структуре. Предоставляет
контекст для этого алгоритма и сохраняет его локальное состояние
!!! Часто в состоянии аккумулируются результаты, полученные в процессе обхода структуры

- Element (Node) - элемент
Определяет операцию Accept, которая принимает посетителя в аргументе

- ConcreteElement (AssignmentNode, VariableRefNode) - конкретный элемент

- ObjectStructure (Program) - структура объектов
Может перечислять свои элементы
Может предоставить посетителю высокоуровневый интерфейс для посещения своих элементов
Может быть как составным объектом так и коллекцией, например списком или множеством

ОТНОШЕНИЯ

- Клиент, использующий паттерн посетитель, должен создать объект класса ConcreteVisitor,
а затем обойти всю структуру, посетив каждый ее элемент

- При посещении элемента последний вызывает операцию посетителя, соответствующую своему классу.
Элемент передает этой операции себя в аргументе, чтобы посетитель мог при необходимости
получить доступ к его состоянию

ДОСТОИНСТВА

- Упрощает добавление новых операций.
Для добавления новой операции над структурой достаточно просто ввести нового посетителя

- Объединение родственных операций и отсечение тех, которые не имеют к ним отношения.
Родственное поведение не разносится по всем классам, присутствующим в структуре объектов,
оно локализировано в посетителе. Не связанные друг с другом функции распределяются по
отдельным подклассам класса Visitor.

- Накопление состояния
Если не использовать этот паттерн, то состояние придется передавать в дополнительных
аргументах операций, выполняющих обход, или хранить в глобальных переменных

НЕДОСТАТКИ

- Если функциональность распределена по нескольким классам, то для определения новой
операции придется изменять каждый класс

- Трудности с добавлением новых классов ConcreteElement.
Паттерн посетитель усложняет добавление новых подклассов класса Element.
Каждый новый конкретный элемент требует объявления новой абстрактной операции в классе
Visitor, которую нужно реализовать в каждом из существующих классов ConcreteVisitor.

- Нарушение инкапсуляции
Применение посетителей подразумевает, что класс ConcreteElement имеет достаточно развитый
интерфейс, для того чтобы посетители могли справиться со своей работой. Поэтому при
использовании данного паттерна приходится предоставлять открытые операции для доступа
к внутреннему состоянию элементов, что ставить под угрозу инкапсуляцию

ДВОЙНАЯ ДИСПЕТЧЕРЗАЦИЯ
Посетитель добавляет в классы новые операции без их изменения.
*/

// Element - интерфейс элемента
type Element interface {
	Accept(visitor Visitor)
}

// ConcreteElementA - представляет конкретный элемент
type ConcreteElementA struct {
	Name string
}

// ConcreteElementB - представляет еще один конкретный элемент
type ConcreteElementB struct {
	Value int
}

// Accept - позволяет посетителю посетить ConcreteElementA
func (e *ConcreteElementA) Accept(visitor Visitor) {
	visitor.VisitConcreteElementA(e)
}

// Accept - позволяет посетителю посетить ConcreteElementB
func (e *ConcreteElementB) Accept(visitor Visitor) {
	visitor.VisitConcreteElementB(e)
}

// Visitor - объявляет метод Visit для каждого типа ConcreteElement
type Visitor interface {
	VisitConcreteElementA(element *ConcreteElementA)
	VisitConcreteElementB(element *ConcreteElementB)
}

// ConcreteVisitor - реализует интерфейс Visitor
type ConcreteVisitor struct{}

// VisitConcreteElementA - реализует операцию посещения для ConcreteElementA
func (v *ConcreteVisitor) VisitConcreteElementA(element *ConcreteElementA) {
	fmt.Printf("Посетитель посещает ConcreteElementA с именем: %s\n", element.Name)
}

// VisitConcreteElementB - реализует операцию посещения для ConcreteElementB
func (v *ConcreteVisitor) VisitConcreteElementB(element *ConcreteElementB) {
	fmt.Printf("Посетитель посещает ConcreteElementB со значением: %d\n", element.Value)
}

func main() {
	visitor := &ConcreteVisitor{}

	elementA := &ConcreteElementA{Name: "TheVisitorName"}
	elementA.Accept(visitor)

	elementB := &ConcreteElementB{Value: 777}
	elementB.Accept(visitor)
}
