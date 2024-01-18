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

// Shape - элемент
type Shape interface {
	accept(Visitor)
}

// Square - конкретный элемент
type Square struct {
	side int
}

func (s *Square) accept(v Visitor) {
	v.visitForSquare(s)
}

// Circle - конкретный элемент
type Circle struct {
	radius int
}

func (c *Circle) accept(v Visitor) {
	v.visitForCircle(c)
}

// Rectangle - конкретный элемент
type Rectangle struct {
	l int
	b int
}

func (t *Rectangle) accept(v Visitor) {
	v.visitForRectangle(t)
}

// Visitor - интерфейс посетителя
type Visitor interface {
	visitForSquare(*Square)
	visitForCircle(*Circle)
	visitForRectangle(*Rectangle)
}

// AreaCalculator - конкретный посетитель, который реализует интерфейс посетителя
type AreaCalculator struct {
	area int
}

func (a *AreaCalculator) visitForSquare(s *Square) {
	// Calculate area for square.
	// Then assign in to the area instance variable.

	fmt.Println("Calculating area for square")
}

func (a *AreaCalculator) visitForCircle(s *Circle) {
	fmt.Println("Calculating area for circle")
}

func (a *AreaCalculator) visitForRectangle(s *Rectangle) {
	fmt.Println("Calculating area for rectangle")
}

// MiddleCoordinates - конкретный посетитель, который реализует интерфейс посетителя
type MiddleCoordinates struct {
	x int
	y int
}

func (a *MiddleCoordinates) visitForSquare(s *Square) {
	// Calculate middle point coordinates for square.
	// Then assign in to the x and y instance variable.
	fmt.Println("Calculating middle point coordinates for square")
}

func (a *MiddleCoordinates) visitForCircle(c *Circle) {
	fmt.Println("Calculating middle point coordinates for circle")
}
func (a *MiddleCoordinates) visitForRectangle(t *Rectangle) {
	fmt.Println("Calculating middle point coordinates for rectangle")
}
func main() {
	square := &Square{side: 2}
	circle := &Circle{radius: 3}
	rectangle := &Rectangle{l: 2, b: 3}

	areaCalculator := &AreaCalculator{}

	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	rectangle.accept(areaCalculator)

	fmt.Println()
	middleCoordinates := &MiddleCoordinates{}
	square.accept(middleCoordinates)
	circle.accept(middleCoordinates)
	rectangle.accept(middleCoordinates)
}
